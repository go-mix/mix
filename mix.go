// Package atomix is a sequence-based Go-native audio mixer
package atomix

import (
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/outrightmental/go-atomix/bind"
)

// Tz is the unit of measurement of samples-over-time, e.g. for 48000Hz playback there are 48,000 Tz in 1 second.
type Tz uint64

/*
 *
 private */

var (
	mixMutex       = &sync.Mutex{}
	mixStartAtTime time.Time
	mixNowTz       Tz
	mixNextCycleTz Tz
	mixCycleTz     Tz
	mixTzDur       time.Duration
	// TODO: implement mixFreq float64
	mixSource       map[string]*Source
	mixSourcePrefix string
	mixReadyFires   []*Fire
	mixLiveFires    []*Fire
	mixSpec         *bind.AudioSpec
	mixFreq         float64
	mixChannels     float64
	isDebug         bool
)

func mixDebug(isOn bool) {
	isDebug = isOn
}

func mixDebugf(format string, args ...interface{}) {
	if isDebug {
		fmt.Printf(format, args...)
	}
}

func mixStartAt(t time.Time) {
	mixStartAtTime = t
}

func mixGetStartTime() time.Time {
	return mixStartAtTime
}

func mixSourceLength(source string) Tz {
	s := mixGetSource(source)
	if s == nil {
		return 0
	}
	return s.Length()
}

func mixSetFire(source string, begin time.Duration, sustain time.Duration, volume float64, pan float64) *Fire {
	mixPrepareSource(mixSourcePrefix + source)
	beginTz := Tz(begin.Nanoseconds() / mixTzDur.Nanoseconds())
	var endTz Tz
	if sustain != 0 {
		endTz = beginTz + Tz(sustain.Nanoseconds()/mixTzDur.Nanoseconds())
	}
	fire := NewFire(mixSourcePrefix+source, beginTz, endTz, volume, pan)
	mixReadyFires = append(mixReadyFires, fire)
	return fire
}

func mixClearAllFires() {
	mixReadyFires = make([]*Fire, 0)
	mixLiveFires = make([]*Fire, 0)
}

func mixSetSoundsPath(prefix string) {
	mixSourcePrefix = prefix
}

func mixSetSpec(s bind.AudioSpec) {
	mixSpec = &s
	mixFreq = float64(s.Freq)
	mixChannels = float64(s.Channels)
	mixTzDur = time.Second / time.Duration(mixFreq)
	mixCycleTz = Tz(mixFreq) // For now, the cycle is always 1 second
}

func mixFireCount() int {
	return len(mixLiveFires) + len(mixReadyFires)
}

func mixTeardown() {
	bind.Teardown()
}

func mixNextSample() []float64 {
	sample := make([]float64, mixSpec.Channels)
	var fireSample []float64
	for _, fire := range mixLiveFires {
		if fireTz := fire.At(mixNowTz); fireTz > 0 {
			fireSample = mixSourceAt(fire.Source, fire.Volume, fire.Pan, fireTz)
			for c := 0; c < mixSpec.Channels; c++ {
				sample[c] += fireSample[c]
			}
		}
	}
	//	mixDebugf("*Mixer.nextSample %+v\n", sample)
	mixNowTz++
	out := make([]float64, mixSpec.Channels)
	for c := 0; c < mixSpec.Channels; c++ {
		out[c] = mixLogarithmicRangeCompression(sample[c])
	}
	if mixNowTz > mixNextCycleTz {
		mixCycle()
	}
	return out
}

func mixSourceAt(src string, volume float64, pan float64, at Tz) []float64 {
	s := mixGetSource(src)
	if s == nil {
		return make([]float64, mixSpec.Channels)
	}
	// if at != 0 {
	// 	mixDebugf("About to source.SampleAt %v in %v\n", at, s.URL)
	// }
	return s.SampleAt(at, volume, pan)
}

func mixPrepareSource(source string) {
	mixMutex.Lock()
	defer mixMutex.Unlock()
	if _, exists := mixSource[source]; !exists {
		mixSource[source] = NewSource(source)
	}
}

func mixGetSource(source string) *Source {
	mixMutex.Lock()
	defer mixMutex.Unlock()
	if _, ok := mixSource[source]; ok {
		return mixSource[source]
	}
	return nil
}

func mixCycle() {
	var fire *Fire
	// if a fire is near-to-playback, move it to the live fire queue
	keepReadyFires := make([]*Fire, 0)
	for _, fire = range mixReadyFires {
		if fire.BeginTz < mixNowTz + mixCycleTz * 2 { // for now, double a mix cycle is consider near-playback
			mixLiveFires = append(mixLiveFires, fire)
		} else {
			keepReadyFires = append(keepReadyFires, fire)
		}
	}
	mixReadyFires = keepReadyFires
	// keep only active fires
	keepLiveFires := make([]*Fire, 0)
	for _, fire = range mixLiveFires {
		if fire.IsAlive() {
			keepLiveFires = append(keepLiveFires, fire)
		} else {
			fire.Teardown()
		}
	}
	mixDebugf("[cycle@%d] ready:%d active:%d\n", mixNowTz, len(mixReadyFires), len(mixLiveFires))
	mixLiveFires = keepLiveFires
	mixNextCycleTz = mixNowTz + mixCycleTz
}

func mixLogarithmicRangeCompression(i float64) float64 {
	if i < -1 {
		return -math.Log(-i-0.85)/14 - 0.75
	} else if i > 1 {
		return math.Log(i-0.85)/14 + 0.75
	} else {
		return i / 1.61803398875
	}
}

// volume (0 to 1), and pan (-1 to +1)
// TODO: implicit panning of source channels! e.g. 2 channels is full left, full right.
func mixVolume(channel float64, volume float64, pan float64) float64 {
	if pan == 0 {
		return volume
	} else if pan < 0 {
		return math.Max(0, 1+pan*channel/mixChannels)
	} else { // pan > 0
		return math.Max(0, 1-pan*channel/mixChannels)
	}
}

func init() {
	mixSource = make(map[string]*Source, 0)
	mixStartAtTime = time.Now().Add(0xFFFF * time.Hour) // this gets reset by Start() or StartAt()
}
