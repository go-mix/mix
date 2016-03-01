// Package ontomix is a sequence-based Go-native audio mixer
package ontomix

import (
	"time"

	"gopkg.in/ontomix.v0/bind"
	"gopkg.in/ontomix.v0/bind/debug"
	"gopkg.in/ontomix.v0/bind/spec"

	"gopkg.in/ontomix.v0/lib/fire"
	"gopkg.in/ontomix.v0/lib/mix"
)

// VERSION # of this ontomix source code
// const VERSION = "0.0.3"

// Debug ON/OFF (ripples down to all sub-modules)
func Debug(isOn bool) {
	debug.Configure(isOn)
}

// Configure the mixer frequency, format, channels & sample rate.
func Configure(s spec.AudioSpec) {
	s.Validate()
	bind.SetOutputCallback(mix.NextSample)
	bind.Configure(s)
	mix.Configure(s)
}

// Teardown everything and release all memory.
func Teardown() {
	mix.Teardown()
	bind.Teardown()
}

// Spec for the mixer, which may include callback functions, e.g. portaudio
func Spec() *spec.AudioSpec {
	return mix.Spec()
}

// SetFire to represent a single audio source playing at a specific time in the future (in time.Duration from play start), with sustain time.Duration, volume from 0 to 1, and pan from -1 to +1
func SetFire(source string, begin time.Duration, sustain time.Duration, volume float64, pan float64) *fire.Fire {
	return mix.SetFire(source, begin, sustain, volume, pan)
}

// FireCount to check the number of fires currently scheduled for playback
func FireCount() int {
	return mix.FireCount()
}

// ClearAllFires to clear all fires currently ready, or live
func ClearAllFires() {
	mix.ClearAllFires()
}

// SetSoundsPath prefix
func SetSoundsPath(prefix string) {
	mix.SetSoundsPath(prefix)
}

// Set the duration between "mix cycles", wherein garbage collection is performed.
func SetMixCycleDuration(d time.Duration) {
	mix.SetCycleDuration(d)
}

// Start the mixer now
func Start() {
	mix.StartAt(time.Now())
}

// StartAt a specific time in the future
func StartAt(t time.Time) {
	mix.StartAt(t)
}

// GetStartTime the mixer was started at
func GetStartTime() time.Time {
	return mix.GetStartTime()
}

// OutputBegin to output WAV opener as []byte via stdout
func OutputBegin() {
	mix.OutputBegin()
}

// OutputBegin to  mix and output as []byte via stdout, up to a specified duration-since-start
func OutputContinueTo(t time.Duration) {
	mix.OutputContinueTo(t)
}

// OutputBegin to output WAV closer as []byte via stdout
func OutputClose() {
	mix.OutputClose()
}
