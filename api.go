// Package atomix is a sequence-based Go-native audio mixer
package atomix

import (
	"time"

	"github.com/outrightmental/go-atomix/bind"
)

// VERSION # of this go-atomix source code
const VERSION = "0.0.2"

// Debug ON/OFF (ripples down to all sub-modules)
func Debug(isOn bool) {
	mixDebug(isOn)
}

// Configure the mixer frequency, format, channels & sample rate.
func Configure(spec bind.AudioSpec) {
	if spec.Freq == 0 {
		panic("Must specify Frequency")
	} else if spec.Format == 0 {
		panic("Must specify Format")
	} else if spec.Channels == 0 {
		panic("Must specify Channels")
	}
	bind.SetMixNextOutput(mixNextSample)
	mixSetSpec(spec)
}

// Teardown everything and release all memory.
func Teardown() {
	mixTeardown()
}

// Cleanup approximately once per second to conserve memory.
func Cleanup() {
	mixCleanup()
}

// Spec for the mixer, which may include callback functions, e.g. go-sdl2
func Spec() *bind.AudioSpec {
	return mixSpec
}

// SetFire to represent a single audio source playing at a specific time in the future (in time.Duration from play start), with sustain time.Duration, volume from 0 to 1, and pan from -1 to +1
func SetFire(source string, begin time.Duration, sustain time.Duration, volume float64, pan float64) *Fire {
	return mixSetFire(source, begin, sustain, volume, pan)
}

// SetSoundsPath prefix
func SetSoundsPath(prefix string) {
	mixSetSoundsPath(prefix)
}

// Start the mixer now
func Start() {
	mixStartAt(time.Now())
}

// StartAt a specific time in the future
func StartAt(t time.Time) {
	mixStartAt(t)
}

// GetStartTime the mixer was started at
func GetStartTime() time.Time {
	return mixGetStartTime()
}

// OpenAudio begins streaming to the bound output audio interface, via a callback function
func OpenAudio() {
	bind.OpenAudio(mixSpec)
}
