// Package bind is for modular binding of atomix to audio interface
package bind

/*
#include <stdio.h>
#include <stdint.h>
typedef unsigned char Uint8;
void AudioCallback(void *userdata, Uint8 *stream, int len);
*/
import "C"
import (
	"reflect"
	"unsafe"
	"github.com/veandco/go-sdl2/sdl"
)

func sdl2OpenAudio(spec *AudioSpec) {
	if err := sdl.Init(sdl.INIT_AUDIO); err != nil {
		panic(err)
	}
	sdl.OpenAudio(sdl2Spec(spec), nil)
	sdl.PauseAudio(false)
}

func sdl2Spec(spec *AudioSpec) *sdl.AudioSpec {
	return &sdl.AudioSpec{
		Freq: spec.Freq,
		Format: sdl2Format(spec.Format),
		Channels: uint8(spec.Channels),
		Samples: 4096,
		Callback: sdl.AudioCallback(C.AudioCallback),
	}
}

func sdl2Unspec(sdlSpec *sdl.AudioSpec) *AudioSpec {
	return &AudioSpec{
		Freq: sdlSpec.Freq,
		Channels: uint16(sdlSpec.Channels),
		Format: sdl2Unformat(sdlSpec.Format),
	}
}

func sdl2Teardown() {
	sdl.PauseAudio(true)
	sdl.Quit()
}

func sdl2Format(fmt AudioFormat) sdl.AudioFormat {
	switch fmt {
	case AudioU8:
		return sdl.AUDIO_U8
	case AudioS8:
		return sdl.AUDIO_S8
	case AudioU16LSB:
		return sdl.AUDIO_U16LSB
	case AudioS16LSB:
		return sdl.AUDIO_S16LSB
	case AudioU16MSB:
		return sdl.AUDIO_U16MSB
	case AudioS16MSB:
		return sdl.AUDIO_S16MSB
	case AudioS32LSB:
		return sdl.AUDIO_S32LSB
	case AudioS32MSB:
		return sdl.AUDIO_S32MSB
	case AudioF32LSB:
		return sdl.AUDIO_F32LSB
	case AudioF32MSB:
		return sdl.AUDIO_F32MSB
	}
	return sdl.AudioFormat(0)
}

func sdl2Unformat(sdlFmt sdl.AudioFormat) AudioFormat {
	switch sdlFmt {
	case sdl.AUDIO_U8:
		return AudioU8
	case sdl.AUDIO_S8:
		return AudioS8
	case sdl.AUDIO_U16LSB:
		return AudioU16LSB
	case sdl.AUDIO_S16LSB:
		return AudioS16LSB
	case sdl.AUDIO_U16MSB:
		return AudioU16MSB
	case sdl.AUDIO_S16MSB:
		return AudioS16MSB
	case sdl.AUDIO_S32LSB:
		return AudioS32LSB
	case sdl.AUDIO_S32MSB:
		return AudioS32MSB
	case sdl.AUDIO_F32LSB:
		return AudioF32LSB
	case sdl.AUDIO_F32MSB:
		return AudioF32MSB
	}
	return AudioFormat(0)
}

// AudioCallback is an unsafe C++ callback function for go-sdl2
//export AudioCallback
func AudioCallback(userdata unsafe.Pointer, stream *C.Uint8, length C.int) {
	byteSize := int(length)
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(stream)),
		Len:  byteSize,
		Cap:  byteSize,
	}
	buf := *(*[]C.Uint8)(unsafe.Pointer(&hdr))

	output := mixNextOutput(byteSize)
	if output == nil {
		// TODO: evaluate whether this failure is productive, or what else could be
		panic("Nil output buffer")
	}
	for i := 0; i < byteSize; i++ {
		buf[i] = C.Uint8(output[i])
	}
}

/*

// DEPRECATED IN FAVOR OF go-wav

func sdl2LoadWAV(file string, spec *AudioSpec) ([]byte, *AudioSpec) {
	data, sdlSpec := sdl.LoadWAV(file, sdl2Spec(spec))
	return data, sdl2Unspec(sdlSpec)
}


func (s *Source) load8(data []byte) {
	channels := int(s.spec.Channels)
	for n := 0; n < len(data); n++ {
		sample := make([]float64, channels)
		for c := 0; c < channels; c++ {
			switch s.spec.Format {
			case bind.AudioU8:
				sample[c] = sampleByteU8(data[n])
			case bind.AudioS8:
				sample[c] = sampleByteS8(data[n])
			}
		}
		s.sample = append(s.sample, sample)
	}
	mixDebugf("*Source[%s].load8(...) length %d channels %d\n", s.URL, len(s.sample), s.spec.Channels)
}

func (s *Source) load16(data []byte) {
	channels := int(s.spec.Channels)
	for n := 0; n < len(data); n += 2 {
		sample := make([]float64, channels)
		for c := 0; c < channels; c++ {
			b := n + c*2
			switch s.spec.Format {
			case bind.AudioU16LSB:
				sample[c] = sampleBytesU16LSB(data[b : b+2])
			case bind.AudioS16LSB:
				sample[c] = sampleBytesS16LSB(data[b : b+2])
			case bind.AudioU16MSB:
				sample[c] = sampleBytesU16MSB(data[b : b+2])
			case bind.AudioS16MSB:
				sample[c] = sampleBytesS16MSB(data[b : b+2])
			}
		}
		s.sample = append(s.sample, sample)
	}
	mixDebugf("*Source[%s].load16(...) length %d channels %d\n", s.URL, len(s.sample), s.spec.Channels)
}

func (s *Source) load32(data []byte) {
	channels := int(s.spec.Channels)
	for n := 0; n < len(data); n += channels * 4 {
		sample := make([]float64, channels)
		for c := 0; c < channels; c++ {
			b := n + c*4
			switch s.spec.Format {
			case bind.AudioS32LSB:
				sample[c] = sampleBytesS32LSB(data[b : b+4])
			case bind.AudioS32MSB:
				sample[c] = sampleBytesS32MSB(data[b : b+4])
			case bind.AudioF32LSB:
				sample[c] = sampleBytesF32LSB(data[b : b+4])
			case bind.AudioF32MSB:
				sample[c] = sampleBytesF32MSB(data[b : b+4])
			}
		}
		s.sample = append(s.sample, sample)
	}
	mixDebugf("*Source[%s].load32(...) length %d channels %d\n", s.URL, len(s.sample), s.spec.Channels)
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
*/
