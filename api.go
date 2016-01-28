// Package atomix is a sequence-based Go-native audio mixer
package atomix

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

const VERSION = "0.0.2"

func Debug(isOn bool) {
	mixDebug(isOn)
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
	mixSetSpec(spec)
}

func Teardown() {
	mixTeardown()
}

func Spec() *sdl.AudioSpec {
	return mixGetSpec()
}

func SetFire(source string, begin time.Duration, sustain time.Duration, volume float64, pan float64) *Fire {
	return mixSetFire(source, begin, sustain, volume, pan)
}

func SetSoundsPath(prefix string) {
	mixSetSoundsPath(prefix)
}

func Start() {
	mixStartAt(time.Now())
}

func StartAt(t time.Time) {
	mixStartAt(t)
}

func GetStartTime() time.Time {
	return mixGetStartTime()
}

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
