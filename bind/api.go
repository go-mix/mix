// Package bind is for modular binding of atomix to audio interface
package bind

// OpenAudio begins streaming to the bound output audio interface, via a callback function
func OpenAudio(spec *AudioSpec) {
	sdl2OpenAudio(spec)
}

// SetMixNextOutputFunc to stream mix output from go-atomix
func SetMixNextOutput(fn mixNextOutputFunc) {
	mixNextOutput = fn
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
	sdl2Teardown()
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
	Freq     int32
	Format   AudioFormat
	Channels uint16
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

// AudioU16LSB is integer unsigned 16-bit, least-significant-bit order
const AudioU16LSB AudioFormat = 16

// AudioS16LSB is integer signed 16-bit, least-significant-bit order
const AudioS16LSB AudioFormat = 17

// AudioU16MSB is integer unsigned 16-bit, most-significant-bit order
const AudioU16MSB AudioFormat = 18

// AudioS16MSB is integer signed 16-bit, most-significant-bit order
const AudioS16MSB AudioFormat = 19

// AudioS32LSB is integer signed 32-bit, least-significant-bit order
const AudioS32LSB AudioFormat = 32

// AudioS32MSB is integer signed 32-bit, most-significant-bit order
const AudioS32MSB AudioFormat = 33

// AudioF32LSB is floating-point 32-bit, least-significant-bit order
const AudioF32LSB AudioFormat = 35

// AudioF32MSB is floating-point 32-bit, most-significant-bit order
const AudioF32MSB AudioFormat = 36

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

type mixNextOutputFunc func (byteSize int) []byte

var (
	mixNextOutput mixNextOutputFunc
	useWAV = OptWAVGo
	usePlayback = OptPlaybackSDL
)
