// Package ontomix is a sequence-based Go-native audio mixer.
//
// Go-native audio mixer for Music apps
//
// See `demo/demo.go`:
//
//     package main
//
//     import (
//       "fmt"
//       "os"
//       "time"
//
//       "github.com/go-ontomix/ontomix"
//       "github.com/go-ontomix/ontomix/bind"
//     )
//
//     var (
//       sampleHz   = float64(48000)
//       spec = bind.AudioSpec{
//         Freq:     sampleHz,
//         Format:   bind.AudioF32,
//         Channels: 2,
//         }
//       bpm        = 120
//       step       = time.Minute / time.Duration(bpm*4)
//       loops      = 16
//       prefix     = "sound/808/"
//       kick1      = "kick1.wav"
//       kick2      = "kick2.wav"
//       marac      = "maracas.wav"
//       snare      = "snare.wav"
//       hitom      = "hightom.wav"
//       clhat      = "cl_hihat.wav"
//       pattern    = []string{
//         kick2,
//         marac,
//         clhat,
//         marac,
//         snare,
//         marac,
//         clhat,
//         kick2,
//         marac,
//         marac,
//         hitom,
//         marac,
//         snare,
//         kick1,
//         clhat,
//         marac,
//       }
//     )
//
//     func main() {
//       defer ontomix.Teardown()
//
//       ontomix.Debug(true)
//       ontomix.Configure(spec)
//       ontomix.SetSoundsPath(prefix)
//       ontomix.StartAt(time.Now().Add(1 * time.Second))
//
//       t := 2 * time.Second // padding before music
//       for n := 0; n < loops; n++ {
//         for s := 0; s < len(pattern); s++ {
//           ontomix.SetFire(pattern[s], t+time.Duration(s)*step, 0, 1.0, 0)
//         }
//         t += time.Duration(len(pattern)) * step
//       }
//
//       fmt.Printf("Ontomix, pid:%v, spec:%v\n", os.Getpid(), spec)
//       for ontomix.FireCount() > 0 {
//         time.Sleep(1 * time.Second)
//       }
//     }
//
// Play this Demo from the root of the project, using SDL 2.0 for hardware audio playback:
//
//     make demo
//
// Or export WAV via stdout `> demo/demo-output.wav`:
//
//     make demo.wav
//
// What
//
// Game audio mixers are designed to play audio spontaneously, but when the timing is known in advance (e.g. sequence-based music apps) there is a demand for much greater accuracy in playback timing.
//
// Read the API documentation at [godoc.org/github.com/go-ontomix/ontomix](https://godoc.org/github.com/go-ontomix/ontomix)
//
// **Ontomix** seeks to solve the problem of audio mixing for the purpose of the playback of sequences where audio files and their playback timing is known in advance.
//
// Ontomix stores and mixes audio in native Go `[]float64` and natively implements Paul Vögler's "Loudness Normalization by Logarithmic Dynamic Range Compression" (details below)
//
// Author: [Charney Kaye](http://w.charney.io)
//
// NOTICE: THIS PROJECT IS IN ALPHA STAGE, AND THE API MAY BE SUBJECT TO CHANGE
//
// Best efforts will be made to preserve each API version in a release tag that can be parsed, e.g. **[github.com/go-ontomix/ontomix](http://github.com/go-ontomix/ontomix)**
//
// Why
//
// Even after selecting a hardware interface library such as [PortAudio](http://www.portaudio.com/) or [C++ SDL 2.0](https://www.libsdl.org/), there remains a critical design problem to be solved.
//
// This design is a **music application mixer**. Most available options are geared towards Game development.
//
// Game audio mixers offer playback timing accuracy +/- 2 milliseconds. But that's totally unacceptable for music, specifically sequence-based sample playback.
//
// The design pattern particular to Game design is that the timing of the audio is not know in advance- the timing that really matters is that which is assembled in near-real-time in response to user interaction.
//
// In the field of Music development, often the timing is known in advance, e.g. a **sequencer**, the composition of music by specifying exactly how, when and which audio files will be played relative to the beginning of playback.
//
// Ergo, **ontomix** seeks to solve the problem of audio mixing for the purpose of the playback of sequences where audio files and their playback timing is known in advance. It seeks to do this with the absolute minimal logical overhead on top of the audio interface.
//
// Ontomix takes maximum advantage of Go by storing and mixing audio in native Go `[]float64` and natively implementing Paul Vögler's "Loudness Normalization by Logarithmic Dynamic Range Compression"
//
// Time
//
// To the Ontomix API, time is specified as a time.Duration-since-epoch, where the epoch is the moment that ontomix.Start() was called.
//
// Internally, time is tracked as samples-since-epoch at the master out playback frequency (e.g. 48000 Hz). This is most efficient because source audio is pre-converted to the master out playback frequency, and all audio maths are performed in terms of samples.
//
// The Mixing Algorithm
//
// Inspired by the theory paper "Mixing two digital audio streams with on the fly Loudness Normalization by Logarithmic Dynamic Range Compression" by Paul Vögler, 2012-04-20. A .PDF has been included [here](docs/LogarithmicDynamicRangeCompression-PaulVogler.pdf), from the paper originally published [here](http://www.voegler.eu/pub/audio/digital-audio-mixing-and-normalization.html).
//
// Usage
//
// There's a demo implementation of **ontomix** included in the `demo/` folder in this repository. Run it using the defaults:
//
//     go run 808.go
//
// Or specify options, e.g. using SDL for playback
//
//     go run 808.go --playback sdl
//
// To show the help screen:
//
//     go run 808.go --help
//
// Dependencies
//
// SDL2
//
// Ubuntu:
//
//     sudo apt-get install libsdl2-dev
//
// Mac OS X:
//
//     brew install sdl2
//
// More details for Linux, Mac and Windows: ["Setting up SDL" by Lazy Foo' Productions](http://lazyfoo.net/SDL_tutorials/lesson01/index.php)
//
// Portaudio
//
// Ubuntu:
//
//     sudo apt-get install portaudio19-dev
//
// Mac OS X:
//
//     brew install portaudio
//
// Windows: ["Building and Installing PortAudio on Windows" by GNU Radio](https://gnuradio.org/redmine/projects/gnuradio/wiki/PortAudioInstall).
//
package ontomix

import (
	"time"

	"github.com/go-ontomix/ontomix/bind"
	"github.com/go-ontomix/ontomix/bind/debug"
	"github.com/go-ontomix/ontomix/bind/spec"

	"github.com/go-ontomix/ontomix/lib/fire"
	"github.com/go-ontomix/ontomix/lib/mix"
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

// OutputBegin to mix and output as []byte via stdout, up to a specified duration-since-start
func OutputContinueTo(t time.Duration) {
	mix.OutputContinueTo(t)
}

// OutputBegin to output WAV closer as []byte via stdout
func OutputClose() {
	mix.OutputClose()
}
