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

// The version # of this go-atomix source code
const VERSION = "0.0.2"

// Turn on debugging (ripples down to all sub-modules)
func Debug(isOn bool) {
	mixDebug(isOn)
}

// Configure the mixer frequency, format, channels & sample rate.
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

// Teardown everything and release all memory.
func Teardown() {
	mixTeardown()
}

// Return the mixer Spec, which may include callback functions, e.g. go-sdl2
func Spec() *sdl.AudioSpec {
	return mixGetSpec()
}

// Set a "Fire" to represent a single audio source playing at a specific time in the future.
func SetFire(source string, begin time.Duration, sustain time.Duration, volume float64, pan float64) *Fire {
	return mixSetFire(source, begin, sustain, volume, pan)
}

// Set the master sounds path prefix
func SetSoundsPath(prefix string) {
	mixSetSoundsPath(prefix)
}

// Start the mixer now
func Start() {
	mixStartAt(time.Now())
}

// Start the mixer at a specified time in the future
func StartAt(t time.Time) {
	mixStartAt(t)
}

// Get the time the mixer was started at
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
