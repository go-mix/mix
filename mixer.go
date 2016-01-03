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
	"fmt"
	"reflect"
	"unsafe"
	"sync"
	// log "github.com/Sirupsen/logrus"
	"github.com/veandco/go-sdl2/sdl"
)


type Mixer struct {
	// nothing yet
}

func (m *Mixer) Initialize() {
	// nothing yet
}

var instance *Mixer
var once sync.Once

func Instance() *Mixer {
	once.Do(func() {
		instance = &Mixer{}
		instance.Initialize()
	})
	return instance
}

func Spec(spec *sdl.AudioSpec) *sdl.AudioSpec {
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
	return spec
}

var	(
	defaultSample = uint8(0xFF)
	defaultAudio = C.Uint8(defaultSample)
)

//export AudioCallback
func AudioCallback(userdata unsafe.Pointer, stream *C.Uint8, length C.int) {
	n := int(length)
	hdr := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(stream)), Len: n, Cap: n}
	buf := *(*[]C.Uint8)(unsafe.Pointer(&hdr))

	for i := 0; i < n; i += 1 {
		buf[i] = defaultAudio
	}
	fmt.Printf("AudioCallback length %d\n", n)
}
