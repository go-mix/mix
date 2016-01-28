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

func sdl2LoadWAV(file string, spec *AudioSpec) ([]byte, *AudioSpec) {
	data, sdlSpec := sdl.LoadWAV(file, sdl2Spec(spec))
	return data, sdl2Unspec(sdlSpec)
}

func sdl2Spec(spec *AudioSpec) *sdl.AudioSpec {
	return &sdl.AudioSpec{
		Freq: spec.Freq,
		Format: sdl2Format(spec.Format),
		Channels: spec.Channels,
		Samples: 4096,
		Callback: sdl.AudioCallback(C.AudioCallback),
	}
}

func sdl2Unspec(sdlSpec *sdl.AudioSpec) *AudioSpec {
	return &AudioSpec{
		Freq: sdlSpec.Freq,
		Channels: sdlSpec.Channels,
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
