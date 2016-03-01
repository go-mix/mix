// Package ontomix is a sequence-based Go-native audio mixer
package ontomix

import (
	"time"

	"gopkg.in/ontomix.v0/bind"
)

// VERSION # of this ontomix source code
const VERSION = "0.0.3"

// Debug ON/OFF (ripples down to all sub-modules)
func Debug(isOn bool) {
	mixDebug(isOn)
}

// Configure the mixer frequency, format, channels & sample rate.
func Configure(spec bind.AudioSpec) {
	spec.Validate()
	bind.SetMixNextSample(mixNextSample)
	mixSetSpec(spec)
	bind.OpenAudio(mixSpec)
}

// Teardown everything and release all memory.
func Teardown() {
	mixTeardown()
}

// Spec for the mixer, which may include callback functions, e.g. portaudio
func Spec() *bind.AudioSpec {
	return mixSpec
}

// SetFire to represent a single audio source playing at a specific time in the future (in time.Duration from play start), with sustain time.Duration, volume from 0 to 1, and pan from -1 to +1
func SetFire(source string, begin time.Duration, sustain time.Duration, volume float64, pan float64) *Fire {
	return mixSetFire(source, begin, sustain, volume, pan)
}

// FireCount to check the number of fires currently scheduled for playback
func FireCount() int {
	return mixFireCount()
}

// ClearAllFires to clear all fires currently ready, or live
func ClearAllFires() {
	mixClearAllFires()
}

// SetSoundsPath prefix
func SetSoundsPath(prefix string) {
	mixSetSoundsPath(prefix)
}

// Set the duration between "mix cycles", wherein garbage collection is performed.
func SetMixCycleDuration(d time.Duration) {
	mixSetCycleDuration(d)
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

// OutputBegin to output WAV opener as []byte via stdout
func OutputBegin() {
	mixOutputBegin()
}

// OutputBegin to  mix and output as []byte via stdout, up to a specified duration-since-start
func OutputContinueTo(t time.Duration) {
	mixOutputContinueTo(t)
}

// OutputBegin to output WAV closer as []byte via stdout
func OutputClose() {
	mixOutputClose()
}


