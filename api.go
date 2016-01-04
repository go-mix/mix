/** Copyright 2015 Outright Mental, Inc. */
package atomix // is for sequence mixing

/*
#include <stdio.h>
#include <stdint.h>
typedef unsigned char Uint8;
void AudioCallback(void *userdata, Uint8 *stream, int len);
*/
import "C"
import (
	// "fmt"
	"github.com/veandco/go-sdl2/sdl"
	"reflect"
	// "sync"
	"time"
	"unsafe"
	// "encoding/binary"
)

func Debug(isOn bool) {
	mixer().Debug(isOn)
}

func Configure(spec sdl.AudioSpec) {
	if spec.Freq == 0 {
		panic("Must specify Frequency")
	} else if spec.Format == 0 {
		panic("Must specify Format")
	} else if spec.Channels == 0 {
		panic("Must specify Channels")
	} else if spec.Samples == 0 {
		panic("Must specify Samples")
	}
	spec.Callback = sdl.AudioCallback(C.AudioCallback)
	mixer().setSpec(spec)
}

func Teardown() {
	mixer().Teardown()
}

func Spec() *sdl.AudioSpec {
	return mixer().getSpec()
}

func Play(source string, begin time.Duration, duration time.Duration, volume float64) {
	mixer().Play(source, begin, duration, volume)
}

//export AudioCallback
func AudioCallback(userdata unsafe.Pointer, stream *C.Uint8, length C.int) {
	n := int(length)
	hdr := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(stream)), Len: n, Cap: n}
	buf := *(*[]C.Uint8)(unsafe.Pointer(&hdr))

	for i, b := range mixer().BufferNext(n) {
		buf[i] = C.Uint8(b)
	}
}

