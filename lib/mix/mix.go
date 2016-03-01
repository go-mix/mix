// Package mix combines sources into an output audio stream
package mix

import (
	"math"
	"time"

	"gopkg.in/ontomix.v0/bind/spec"

	"gopkg.in/ontomix.v0/bind/debug"
	"gopkg.in/ontomix.v0/lib/fire"
	"gopkg.in/ontomix.v0/lib/source"
)

// NextSample returns the next sample mixed in all channels
func NextSample() []float64 {
	sample := make([]float64, masterSpec.Channels)
	var fireSample []float64
	for _, fire := range mixLiveFires {
		if fireTz := fire.At(mixNowTz); fireTz > 0 {
			fireSample = mixSourceAt(fire.Source, fire.Volume, fire.Pan, fireTz)
			for c := 0; c < masterSpec.Channels; c++ {
				sample[c] += fireSample[c]
			}
		}
	}
	//	debug.Printf("*Mixer.nextSample %+v\n", sample)
	mixNowTz++
	out := make([]float64, masterSpec.Channels)
	for c := 0; c < masterSpec.Channels; c++ {
		out[c] = mixLogarithmicRangeCompression(sample[c])
	}
	if mixNowTz > mixNextCycleTz {
		mixCycle()
	}
	return out
}

// Configure the mixer frequency, format, channels & sample rate.
func Configure(s spec.AudioSpec) {
	masterSpec = &s
	masterFreq = float64(s.Freq)
	masterTzDur = time.Second / time.Duration(masterFreq)
	masterCycleDurTz = spec.Tz(masterFreq)
	source.Configure(s)
}

// Spec spec returns the current audio specification.
func Spec() *spec.AudioSpec {
	return masterSpec
}

// Teardown everything and release all memory.
func Teardown() {
}

// SetFire to represent a single audio source playing at a specific time in the future (in time.Duration from play start), with sustain time.Duration, volume from 0 to 1, and pan from -1 to +1
func SetFire(source string, begin time.Duration, sustain time.Duration, volume float64, pan float64) *fire.Fire {
	mixPrepareSource(mixSourcePrefix + source)
	beginTz := spec.Tz(begin.Nanoseconds() / masterTzDur.Nanoseconds())
	var endTz spec.Tz
	if sustain != 0 {
		endTz = beginTz + spec.Tz(sustain.Nanoseconds()/masterTzDur.Nanoseconds())
	}
	f := fire.New(mixSourcePrefix+source, beginTz, endTz, volume, pan)
	mixReadyFires = append(mixReadyFires, f)
	return f
}

// FireCount returns the current total ready fires + live fires.
func FireCount() int {
	return len(mixLiveFires) + len(mixReadyFires)
}

// StartAt to specify what time to begin mixing.
func StartAt(t time.Time) {
	mixStartAtTime = t
}

// GetStartTime returns the time mixing began.
func GetStartTime() time.Time {
	return mixStartAtTime
}

// ClearAllFires to remove all ready & live fires.
func ClearAllFires() {
	mixReadyFires = make([]*fire.Fire, 0)
	mixLiveFires = make([]*fire.Fire, 0)
}

// SetSoundsPath to set the sound path prefix.
func SetSoundsPath(prefix string) {
	mixSourcePrefix = prefix
}

// GetCycleDurationTz sets the duration of a mix cycle.
func SetCycleDuration(d time.Duration) {
	if masterFreq == 0 {
		panic("Must specify mixing frequency before setting cycle duration!")
	}
	masterCycleDurTz = spec.Tz((d / time.Second) * time.Duration(masterFreq))
}

// GetCycleDurationTz returns the duration of a mix cycle.
func GetCycleDurationTz() spec.Tz {
	return masterCycleDurTz
}

// OutputBegin to output WAV opener as []byte via stdout
func OutputBegin() {

}

// OutputBegin to  mix and output as []byte via stdout, up to a specified duration-since-start
func OutputContinueTo(t time.Duration) {

}

// OutputBegin to output WAV closer as []byte via stdout
func OutputClose() {

}

/*
 *
 private */

var (
	mixStartAtTime   time.Time
	mixNowTz         spec.Tz
	mixNextCycleTz   spec.Tz
	masterCycleDurTz spec.Tz
	masterTzDur      time.Duration
	// TODO: implement mixFreq float64
	mixSourcePrefix string
	mixReadyFires   []*fire.Fire
	mixLiveFires    []*fire.Fire
	masterSpec      *spec.AudioSpec
	masterFreq      float64
)

func init() {
	mixStartAtTime = time.Now().Add(0xFFFF * time.Hour) // this gets reset by Start() or StartAt()
}

func mixSourceAt(src string, volume float64, pan float64, at spec.Tz) []float64 {
	s := mixGetSource(src)
	if s == nil {
		return make([]float64, masterSpec.Channels)
	}
	// if at != 0 {
	// 	debug.Printf("About to source.SampleAt %v in %v\n", at, s.URL)
	// }
	return s.SampleAt(at, volume, pan)
}

func mixPrepareSource(src string) {
	source.Prepare(src)
}

func mixGetSource(src string) *source.Source {
	return source.Get(src)
}

// TODO:
// Make a new empty map[string]*Source, e.g. keepSource
// While iterating over the ready & active fires (see issues #11 and #18; implemented as of pull #29) copy any used *Source to the new keepSource
// Replace the mixSource with keepSource
func mixCycle() {
	var f *fire.Fire
	// for garbage collection of unused sources:
	keepSource := make(map[string]bool)
	// if a fire is near-to-playback, move it to the live fire queue
	keepReadyFires := make([]*fire.Fire, 0)
	for _, f = range mixReadyFires {
		keepSource[f.Source] = true
		if f.BeginTz < mixNowTz+masterCycleDurTz*2 { // for now, double a mix cycle is consider near-playback
			mixLiveFires = append(mixLiveFires, f)
		} else {
			keepReadyFires = append(keepReadyFires, f)
		}
	}
	mixReadyFires = keepReadyFires
	// keep only active fires
	keepLiveFires := make([]*fire.Fire, 0)
	for _, f = range mixLiveFires {
		if f.IsAlive() {
			keepSource[f.Source] = true
			keepLiveFires = append(keepLiveFires, f)
		} else {
			f.Teardown()
		}
	}
	mixLiveFires = keepLiveFires
	source.Prune(keepSource)
	mixNextCycleTz = mixNowTz + masterCycleDurTz
	if debug.Active() && source.Count() > 0 {
		debug.Printf("ontomix [%dz] fire-ready:%d fire-active:%d sources:%d\n", mixNowTz, len(mixReadyFires), len(mixLiveFires), source.Count())
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
