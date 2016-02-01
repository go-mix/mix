// Package bind is for modular binding of atomix to audio interface
package bind

// OpenAudio begins streaming to the bound output audio interface, via a callback function
func OpenAudio(spec *AudioSpec) {
	portaudioSetup(spec)
}

// SetMixNextOutputFunc to stream mix output from go-atomix
func SetMixNextOutput(fn mixNextOutputFunc) {
	mixNextOutputSample = fn
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
	portaudioTeardown()
}

// UseWAV to select the WAV file interface
func UseWAV(opt OptWAV) {
	useWAV = opt
}

// Use to select the playback interface
func UsePlayback(opt OptPlayback) {
	usePlayback = opt
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

// AudioU8 is integer unsigned 8-bit
const AudioU8 AudioFormat = 1

// AudioS8 is integer signed 8-bit
const AudioS8 AudioFormat = 2

// AudioU16 is integer unsigned 16-bit
const AudioU16 AudioFormat = 16

// AudioS16 is integer signed 16-bit
const AudioS16 AudioFormat = 17

// AudioF32 is floating-point 32-bit
const AudioF32 AudioFormat = 32

// AudioF64 is floating-point 64-bit
const AudioF64 AudioFormat = 64

// OptWAV represents a WAV I/O option
type OptWAV uint8

// OptWAVGo to use Go-Native WAV file I/O
const OptWAVGo OptWAV = 1

// OptPlayback represents a WAV I/O option
type OptPlayback uint8

// OptPlaybackPortAudio to use Go-Native WAV file I/O
const OptPlaybackPortAudio OptPlayback = 1

// OptPlaybackSDL to use SDL for WAV file I/O
const OptPlaybackSDL OptPlayback = 2

/*
 *
 private below here */

type mixNextOutputFunc func () []float64

var (
	mixNextOutputSample mixNextOutputFunc
	useWAV = OptWAVGo
	usePlayback = OptPlaybackSDL
)

func noErr(err error) {
	if err != nil {
		panic(err)
	}
}
