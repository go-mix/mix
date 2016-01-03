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



/*

// read audio file

func ReadAudio() {

	file := "assets/sounds/percussion/808/kick1.wav"

	data, spec := sdl.LoadWAV(file, &sdl.AudioSpec{})

	log.WithFields(log.Fields{
		"spec":   spec,
	}).Info("Loaded")

	for n := 0; n < len(data); n += 2 {
		StoreSample(data[n:n+2])
	}
}

func StoreSample(s []byte) {
	storedAudio = append(storedAudio, C.Uint16(int32(s[0]) + int32(s[1])<<8))
}

var storedAudio []C.Uint16
var	(
	defaultSample = uint16(0xFFFF)
	defaultAudio = C.Uint16(defaultSample)
)

//export AudioCallback
func AudioCallback(userdata unsafe.Pointer, stream *C.Uint16, length C.int) {
	n := int(length)
	hdr := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(stream)), Len: n, Cap: n}
	buf := *(*[]C.Uint16)(unsafe.Pointer(&hdr))

	for i := 0; i < n; i += 1 {
		if i < len(storedAudio) {
			buf[i] = storedAudio[i]
		} else {
			buf[i] = defaultAudio
		}
	}
	fmt.Printf("AudioCallback length %d\n", n)
}

*/
