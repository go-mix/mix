// Package bind is for modular binding of atomix to audio interface
package bind

import (
	"encoding/binary"
	"math"
)

// OpenAudio begins streaming to the bound output audio interface, via a callback function
func OpenAudio(spec *AudioSpec) {
	outputSpec = spec
	switch usePlayback {
	case OptPlaybackPortaudio:
		playPortaudioSetup(spec)
	case OptPlaybackSDL:
		playSDLSetup(spec)
	}
}

// SetMixNextOutputFunc to stream mix output from go-atomix
func SetMixNextSample(fn outputCallbackMixNextSampleFunc) {
	outputCallbackMixNextSample = fn
}

// LoadWAV into a buffer
func LoadWAV(file string) ([][]float64, *AudioSpec) {
	switch useLoader {
	case OptLoaderWAV:
		return LoadNewWAV(file)
	default:
		return make([][]float64, 0), &AudioSpec{}
	}
}

// Teardown to close all hardware bindings
func Teardown() {
	switch usePlayback {
	case OptPlaybackPortaudio:
		playPortaudioTeardown()
	case OptPlaybackSDL:
		playSDLTeardown()
	}
}

// UseLoader to select the file loading interface
func UseLoader(opt string) {
	switch opt {
	case string(OptLoaderWAV):
		useLoader = OptLoaderWAV
	default:
		panic("No such Loader: " + opt)
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

// AudioFormat represents the bit allocation for a single sample of audio
type AudioFormat string

// AudioU8 is unsigned-integer 8-bit sample (per channel)
const AudioU8 AudioFormat = "U8"

// AudioS8 is signed-integer 8-bit sample (per channel)
const AudioS8 AudioFormat = "S8"

// AudioU16 is unsigned-integer 16-bit sample (per channel)
const AudioU16 AudioFormat = "U16"

// AudioS16 is signed-integer 16-bit sample (per channel)
const AudioS16 AudioFormat = "S16"

// AudioS32 is signed-integer 32-bit sample (per channel)
const AudioS32 AudioFormat = "S32"

// AudioF32 is floating-point 32-bit sample (per channel)
const AudioF32 AudioFormat = "F32"

// AudioF64 is floating-point 64-bit sample (per channel)
const AudioF64 AudioFormat = "F64"

// OptLoader represents a WAV I/O option
type OptLoader string

// OptLoadWav to use Go-Native WAV file I/O
const OptLoaderWAV OptLoader = "wav"

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

func sampleByteU8(sample byte) float64 {
	return float64(int8(sample))/float64(0x7F) - float64(1)
}

func sampleByteS8(sample byte) float64 {
	return float64(int8(sample)) / float64(0x7F)
}

func sampleBytesU16LSB(sample []byte) float64 {
	return float64(binary.LittleEndian.Uint16(sample))/float64(0x8000) - float64(1)
}

func sampleBytesU16MSB(sample []byte) float64 {
	return float64(binary.BigEndian.Uint16(sample))/float64(0x8000) - float64(1)
}

func sampleBytesS16LSB(sample []byte) float64 {
	return float64(int16(binary.LittleEndian.Uint16(sample))) / float64(0x7FFF)
}

func sampleBytesS16MSB(sample []byte) float64 {
	return float64(int16(binary.BigEndian.Uint16(sample))) / float64(0x7FFF)
}

func sampleBytesS32LSB(sample []byte) float64 {
	return float64(int32(binary.LittleEndian.Uint32(sample))) / float64(0x7FFFFFFF)
}

func sampleBytesS32MSB(sample []byte) float64 {
	return float64(int32(binary.BigEndian.Uint32(sample))) / float64(0x7FFFFFFF)
}

func sampleBytesF32LSB(sample []byte) float64 {
	return float64(math.Float32frombits(binary.LittleEndian.Uint32(sample)))
}

func sampleBytesF32MSB(sample []byte) float64 {
	return float64(math.Float32frombits(binary.BigEndian.Uint32(sample)))
}

func sampleBytesF64LSB(sample []byte) float64 {
	return float64(math.Float64frombits(binary.LittleEndian.Uint64(sample)))
}

func sampleBytesF64MSB(sample []byte) float64 {
	return float64(math.Float64frombits(binary.BigEndian.Uint64(sample)))
}
