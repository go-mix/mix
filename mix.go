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
	mixTzDur       time.Duration
	// TODO: implement mixFreq float64
	mixSource       map[string]*Source
	mixSourcePrefix string
	mixFires        []*Fire
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
	mixFires = append(mixFires, fire)
	return fire
}

func mixSetSoundsPath(prefix string) {
	mixSourcePrefix = prefix
}

func mixSetSpec(s bind.AudioSpec) {
	mixSpec = &s
	mixFreq = float64(s.Freq)
	mixChannels = float64(s.Channels)
	mixTzDur = time.Second / time.Duration(mixFreq)
}

func mixTeardown() {
	bind.Teardown()
}

func mixNextSample() []float64 {
	sample := make([]float64, mixSpec.Channels)
	var fireSample []float64
	// TODO: #FIXME need a more efficient method of iterating active fires; range fires hogs CPU with >100 fires
	// TODO: #FIXME ^ really this is a serious processor bottleneck. Find a method to avoid iterating over all these inactive fires every sample!
	for _, fire := range mixFires {
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

func mixCleanup() {
	for i, fire := range mixFires {
		if !fire.IsAlive() {
			fire.Teardown()
			mixFires = append(mixFires[:i], mixFires[i+1:]...)
		}
	}
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
