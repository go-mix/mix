// Package mix combines sources into an output audio stream
package mix

import (
	"math"
	"time"

	"github.com/go-ontomix/ontomix/bind/spec"

	"github.com/go-ontomix/ontomix/bind"
	"github.com/go-ontomix/ontomix/bind/debug"
	"github.com/go-ontomix/ontomix/bind/sample"
	"github.com/go-ontomix/ontomix/lib/fire"
	"github.com/go-ontomix/ontomix/lib/source"
)

// NextSample returns the next sample mixed in all channels
func NextSample() []sample.Value {
	smp := make([]sample.Value, masterSpec.Channels)
	var fireSample []sample.Value
	for _, fire := range mixLiveFires {
		if fireTz := fire.At(nowTz); fireTz > 0 {
			fireSample = mixSourceAt(fire.Source, fire.Volume, fire.Pan, fireTz)
			for c := 0; c < masterSpec.Channels; c++ {
				smp[c] += fireSample[c]
			}
		}
	}
	//	debug.Printf("*Mixer.nextSample %+v\n", sample)
	nowTz++
	out := make([]sample.Value, masterSpec.Channels)
	for c := 0; c < masterSpec.Channels; c++ {
		out[c] = mixLogarithmicRangeCompression(smp[c])
	}
	if nowTz > nextCycleTz {
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
	startAtTime = t
}

// GetStartTime returns the time mixing began.
func GetStartTime() time.Time {
	return startAtTime
}

// GetNowAt returns current mix position
func GetNowAt() time.Duration {
	return time.Duration(nowTz) * masterTzDur
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

// OutputStart requires a known length
func OutputStart(length time.Duration) {
	bind.OutputStart(length)
}

// OutputContinueTo to  mix and output as []byte via stdout, up to a specified duration-since-start
func OutputContinueTo(t time.Duration) {
	deltaDur := t - outputToDur
	deltaTz := spec.Tz(masterFreq*float64((deltaDur)/time.Second))
	debug.Printf("mix.OutputContinueTo(%+v) deltaDur:%+v nowTz:%+v deltaTz:%+v begin...", t, deltaDur, nowTz, deltaTz)
	bind.OutputNext(deltaTz)
	outputToDur = t
	debug.Printf("mix.OutputContinueTo(%+v) ...done! nowTz:%+v outputToDur:%+v", t, nowTz, outputToDur)
}

// OutputBegin to output WAV closer as []byte via stdout
func OutputClose() {
	// nothing to do
}

/*
 *
 private */

var (
	outputToDur      time.Duration
	startAtTime      time.Time
	nowTz            spec.Tz
	nextCycleTz      spec.Tz
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
	startAtTime = time.Now().Add(0xFFFF * time.Hour) // this gets reset by Start() or StartAt()
}

func mixSourceAt(src string, volume float64, pan float64, at spec.Tz) []sample.Value {
	s := mixGetSource(src)
	if s == nil {
		return make([]sample.Value, masterSpec.Channels)
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
		if f.BeginTz < nowTz+masterCycleDurTz*2 { // for now, double a mix cycle is consider near-playback
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
	nextCycleTz = nowTz + masterCycleDurTz
	if debug.Active() && source.Count() > 0 {
		debug.Printf("ontomix [%dz] fire-ready:%d fire-active:%d sources:%d\n", nowTz, len(mixReadyFires), len(mixLiveFires), source.Count())
	}
}

func mixLogarithmicRangeCompression(i sample.Value) sample.Value {
	if i < -1 {
		return sample.Value(-math.Log(-float64(i)-0.85)/14 - 0.75)
	} else if i > 1 {
		return sample.Value(math.Log(float64(i)-0.85)/14 + 0.75)
	} else {
		return sample.Value(i / 1.61803398875)
	}
}
