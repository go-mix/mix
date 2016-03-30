// Package bind is for modular binding of mix to audio interface
package bind

import (
	"time"

	"github.com/go-mix/mix/bind/hardware/null"
	"github.com/go-mix/mix/bind/hardware/portaudio"
	"github.com/go-mix/mix/bind/hardware/sdl"
	"github.com/go-mix/mix/bind/opt"
	"github.com/go-mix/mix/bind/sample"
	"github.com/go-mix/mix/bind/spec"
	"github.com/go-mix/mix/bind/wav"
)

// Configure begins streaming to the bound out audio interface, via a callback function
func Configure(s spec.AudioSpec) {
	sample.ConfigureOutput(s)
	switch useOutput {
	case opt.OutputPortAudio:
		portaudio.ConfigureOutput(s)
	case opt.OutputSDL:
		sdl.ConfigureOutput(s)
	case opt.OutputWAV:
		wav.ConfigureOutput(s)
	case opt.OutputNull:
		null.ConfigureOutput(s)
	}
}

func IsDirectOutput() bool {
	return useOutput == opt.OutputWAV
}

// SetMixNextOutFunc to stream mix out from mix
func SetOutputCallback(fn sample.OutNextCallbackFunc) {
	sample.SetOutputCallback(fn)
}

// OutputStart requires a known length
func OutputStart(length time.Duration) {
	switch useOutput {
	case opt.OutputWAV:
		wav.OutputStart(length)
	case opt.OutputNull:
	// do nothing
	}
}

// OutputNext using the configured writer.
func OutputNext(numSamples spec.Tz) {
	switch useOutput {
	case opt.OutputWAV:
		wav.OutputNext(numSamples)
	case opt.OutputNull:
		// do nothing
	}
}

// LoadWAV into a buffer
func LoadWAV(file string) ([]sample.Sample, *spec.AudioSpec) {
	switch useLoader {
	case opt.InputWAV:
		return wav.Load(file)
	default:
		return make([]sample.Sample, 0), &spec.AudioSpec{}
	}
}

// Teardown to close all hardware bindings
func Teardown() {
	switch useOutput {
	case opt.OutputPortAudio:
		portaudio.TeardownOutput()
	case opt.OutputSDL:
		sdl.TeardownOutput()
	case opt.OutputWAV:
		wav.TeardownOutput()
	case opt.OutputNull:
		// do nothing
	}
}

// UseLoader to select the file loading interface
func UseLoader(opt opt.Input) {
	useLoader = opt
}

// UseLoaderString to select the file loading interface by string
func UseLoaderString(loader string) {
	switch loader {
	case string(opt.InputWAV):
		useLoader = opt.InputWAV
	default:
		panic("No such Loader: " + loader)
	}
}

// UseOutput to select the outback interface
func UseOutput(opt opt.Output) {
	useOutput = opt
}

// UseOutputString to select the outback interface by string
func UseOutputString(output string) {
	switch output {
	case string(opt.OutputPortAudio):
		useOutput = opt.OutputPortAudio
	case string(opt.OutputSDL):
		useOutput = opt.OutputSDL
	case string(opt.OutputWAV):
		useOutput = opt.OutputWAV
	case string(opt.OutputNull):
		useOutput = opt.OutputNull
	default:
		panic("No such Output: " + output)
	}
}

/*
 *
 private */

var (
	useLoader = opt.InputWAV
	useOutput = opt.OutputSDL
)
