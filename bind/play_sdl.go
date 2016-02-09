// Package bind is for modular binding of atomix to audio interface
package bind

/*
#include <stdio.h>
#include <stdint.h>
typedef unsigned char Uint8;
void PlaySDLAudioCallback(void *userdata, Uint8 *stream, int len);
*/
import "C"
import (
	"github.com/veandco/go-sdl2/sdl"
	"reflect"
	"unsafe"
)

func playSDLSetup(spec *AudioSpec) {
	if err := sdl.Init(sdl.INIT_AUDIO); err != nil {
		panic(err)
	}
	sdl.OpenAudio(playSDLSpec(), nil)
	sdl.PauseAudio(false)
}

func playSDLSpec() *sdl.AudioSpec {
	return &sdl.AudioSpec{
		Freq:     int32(outputSpec.Freq),
		Format:   playSDLFormat(outputSpec.Format),
		Channels: uint8(outputSpec.Channels),
		Samples:  uint16(1024 * outputSpec.Channels),
		Callback: sdl.AudioCallback(C.PlaySDLAudioCallback),
	}
}

func playSDLTeardown() {
	sdl.PauseAudio(true)
	sdl.Quit()
}

func playSDLFormat(fmt AudioFormat) sdl.AudioFormat {
	switch fmt {
	case AudioU8:
		return sdl.AUDIO_U8
	case AudioS8:
		return sdl.AUDIO_S8
	case AudioU16:
		return sdl.AUDIO_U16
	case AudioS16:
		return sdl.AUDIO_S16
	case AudioF32:
		return sdl.AUDIO_F32
	}
	return sdl.AudioFormat(0)
}

func playSDLNextOutput(byteSize int) (out []byte) {
	for len(out) < byteSize {
		out = append(out, outputNextBytes()...)
	}
	return
}

// PlaySDLAudioCallback is an unsafe C++ callback function for go-sdl2
//export PlaySDLAudioCallback
func PlaySDLAudioCallback(userdata unsafe.Pointer, stream *C.Uint8, length C.int) {
	byteSize := int(length)
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(stream)),
		Len:  byteSize,
		Cap:  byteSize,
	}
	buf := *(*[]C.Uint8)(unsafe.Pointer(&hdr))
	output := playSDLNextOutput(byteSize)
	if output == nil {
		// TODO: evaluate whether this failure is productive, or what else could be
		panic("Nil output buffer")
	}
	for i := 0; i < byteSize; i++ {
		buf[i] = C.Uint8(output[i])
	}
}
