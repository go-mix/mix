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
func LoadWAV(file string, spec *AudioSpec) ([]byte, *AudioSpec) {
	return sdl2LoadWAV(file, spec)
}

// Teardown to close all hardware bindings
func Teardown() {
	sdl2Teardown()
}

// AudioSpec represents the frequency, format, # channels and sample rate of any audio I/O
type AudioSpec struct {
	Freq     int32
	Format   AudioFormat
	Channels uint8
}

// AudioFormat represents the bit allocation for a single sample of audio
type AudioFormat uint8

const (
	// AudioU8 is integer unsigned 8-bit
	AudioU8 AudioFormat = 1

	// AudioS8 is integer signed 8-bit
	AudioS8 AudioFormat = 2

	// AudioU16LSB is integer unsigned 16-bit, least-significant-bit order
	AudioU16LSB AudioFormat = 16

	// AudioS16LSB is integer signed 16-bit, least-significant-bit order
	AudioS16LSB AudioFormat = 17

	// AudioU16MSB is integer unsigned 16-bit, most-significant-bit order
	AudioU16MSB AudioFormat = 18

	// AudioS16MSB is integer signed 16-bit, most-significant-bit order
	AudioS16MSB AudioFormat = 19

	// AudioS32LSB is integer signed 32-bit, least-significant-bit order
	AudioS32LSB AudioFormat = 32

	// AudioS32MSB is integer signed 32-bit, most-significant-bit order
	AudioS32MSB AudioFormat = 33

	// AudioF32LSB is floating-point 32-bit, least-significant-bit order
	AudioF32LSB AudioFormat = 35

	// AudioF32MSB is floating-point 32-bit, most-significant-bit order
	AudioF32MSB AudioFormat = 36
)

/*
 *
 private below here */

type mixNextOutputFunc func (byteSize int) []byte

var mixNextOutput mixNextOutputFunc
