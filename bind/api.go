// Package bind is for modular binding of atomix to audio interface
package bind

// OpenAudio begins streaming to the bound output audio interface, via a callback function
func OpenAudio(spec *AudioSpec) {
	outputSpec = spec
	switch usePlayback {
	case OptPlaybackPortaudio:
		portaudioSetup(spec)
	case OptPlaybackSDL:
		sdlSetup(spec)
	}
}

// SetMixNextOutputFunc to stream mix output from go-atomix
func SetMixNextSample(fn outputCallbackMixNextSampleFunc) {
	outputCallbackMixNextSample = fn
}

// LoadWAV into a buffer
func LoadWAV(file string) ([][]float64, *AudioSpec) {
	switch useWAV {
	case OptWAVGo:
		return nativeLoadWAV(file)
	default:
		return make([][]float64, 0), &AudioSpec{}
	}
}

// Teardown to close all hardware bindings
func Teardown() {
	switch usePlayback {
	case OptPlaybackPortaudio:
		portaudioTeardown()
	case OptPlaybackSDL:
		sdlTeardown()
	}
}

// UseWAV to select the WAV file interface
func UseWAV(opt string) {
	switch opt {
	case string(OptWAVGo):
		useWAV = OptWAVGo
	default:
		panic("No such WAV: " + opt)
	}
}

// Use to select the playback interface
func UsePlayback(opt string) {
	switch opt {
	case string(OptPlaybackPortaudio):
		usePlayback = OptPlaybackPortaudio
	case string(OptPlaybackSDL):
		usePlayback = OptPlaybackSDL
	default:
		panic("No such Playback: " + opt)
	}
}

// AudioSpec represents the frequency, format, # channels and sample rate of any audio I/O
type AudioSpec struct {
	Freq     float64
	Format   AudioFormat
	Channels int
}

//type WavFormat struct {
//	AudioFormat   uint16
//	NumChannels   uint16
//	SampleRate    uint32
//	ByteRate      uint32
//	BlockAlign    uint16
//	BitsPerSample uint16
//}

// AudioFormat represents the bit allocation for a single sample of audio
type AudioFormat uint8

// AudioU8 is unsigned-integer 8-bit sample (per channel)
const AudioU8 AudioFormat = 1

// AudioS8 is signed-integer 8-bit sample (per channel)
const AudioS8 AudioFormat = 2

// AudioU16 is unsigned-integer 16-bit sample (per channel)
const AudioU16 AudioFormat = 16

// AudioS16 is signed-integer 16-bit sample (per channel)
const AudioS16 AudioFormat = 17

// AudioS32 is signed-integer 32-bit sample (per channel)
const AudioS32 AudioFormat = 32

// AudioF32 is floating-point 32-bit sample (per channel)
const AudioF32 AudioFormat = 33

// AudioF64 is floating-point 64-bit sample (per channel)
const AudioF64 AudioFormat = 64

// OptWAV represents a WAV I/O option
type OptWAV string

// OptWAVGo to use Go-Native WAV file I/O
const OptWAVGo OptWAV = "native"

// OptPlayback represents a WAV I/O option
type OptPlayback string

// OptPlaybackPortAudio to use Go-Native WAV file I/O
const OptPlaybackPortaudio OptPlayback = "portaudio"

// OptPlaybackSDL to use SDL for WAV file I/O
const OptPlaybackSDL OptPlayback = "sdl"

/*
 *
 private below here */

func noErr(err error) {
	if err != nil {
		panic(err)
	}
}
