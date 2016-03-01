// Package sdl is for modular binding of ontomix to audio interface via SDL2
package sdl

/*
#include <stdio.h>
#include <stdint.h>
typedef unsigned char Uint8;
void OutSDLAudioCallback(void *userdata, Uint8 *stream, int len);
*/
import "C"
import (
	"reflect"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"

	"gopkg.in/ontomix.v0/bind/sample"
	"gopkg.in/ontomix.v0/bind/spec"
)

func ConfigureOutput(s spec.AudioSpec) {
	outSpec = &s
	if err := sdl.Init(sdl.INIT_AUDIO); err != nil {
		panic(err)
	}
	sdl.OpenAudio(outSDLSpec(), nil)
	sdl.PauseAudio(false)
}

// OutSDLAudioCallback is an unsafe C++ callback function for go-sdl2
//export OutSDLAudioCallback
func OutSDLAudioCallback(userdata unsafe.Pointer, stream *C.Uint8, length C.int) {
	byteSize := int(length)
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(stream)),
		Len:  byteSize,
		Cap:  byteSize,
	}
	buf := *(*[]C.Uint8)(unsafe.Pointer(&hdr))
	out := outSDLNextOut(byteSize)
	if out == nil {
		// TODO: evaluate whether this failure is productive, or what else could be
		panic("Nil out buffer")
	}
	for i := 0; i < byteSize; i++ {
		buf[i] = C.Uint8(out[i])
	}
}

/*
 *
 private */

var (
	outSpec *spec.AudioSpec
)

func outSDLSpec() *sdl.AudioSpec {
	return &sdl.AudioSpec{
		Freq:     int32(outSpec.Freq),
		Format:   outSDLFormat(outSpec.Format),
		Channels: uint8(outSpec.Channels),
		Samples:  uint16(1024 * outSpec.Channels),
		Callback: sdl.AudioCallback(C.OutSDLAudioCallback),
	}
}

func TeardownOutput() {
	sdl.PauseAudio(true)
	sdl.Quit()
}

func outSDLFormat(fmt spec.AudioFormat) sdl.AudioFormat {
	switch fmt {
	case spec.AudioU8:
		return sdl.AUDIO_U8
	case spec.AudioS8:
		return sdl.AUDIO_S8
	case spec.AudioU16:
		return sdl.AUDIO_U16
	case spec.AudioS16:
		return sdl.AUDIO_S16
	case spec.AudioF32:
		return sdl.AUDIO_F32
	}
	return sdl.AudioFormat(0)
}

func outSDLNextOut(byteSize int) (out []byte) {
	for len(out) < byteSize {
		out = append(out, sample.OutNextBytes()...)
	}
	return
}
