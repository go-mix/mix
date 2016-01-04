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
	"github.com/veandco/go-sdl2/sdl"
	// "reflect"
	"sync"
	"time"
	// "unsafe"
	"encoding/binary"
)

var (
	defaultAudio  = C.Uint8(0)
)

// singleton
func mixer() *Mixer {
	once.Do(func() {
		instance = &Mixer{}
		instance.Initialize()
	})
	return instance
}

var (
	instance *Mixer
	once     sync.Once
)

type (
	Hz    uint64
	smp16 uint16
)

type Mixer struct {
	nowAtHz   Hz // current sample since init
	freq      float64 // cache this for maths
	source    map[string]*Source
	fire      []*Fire
	spec      sdl.AudioSpec
	isDebug   bool
}

func (m *Mixer) Initialize() {
	m.source = make(map[string]*Source,0)
	m.nowAtHz = 0
}

func (m *Mixer) Debug(isOn bool) {
	m.isDebug = isOn
}

func (m *Mixer) Printf(format string, args... interface{}) {
	if m.isDebug {
		fmt.Printf(format, args...)
	}
}

func (m *Mixer) Play(source string, begin time.Duration, duration time.Duration, volume float64) {
	m.prepareSource(source)
	m.fire = append(m.fire, NewFire(source, begin, duration, volume))
}

func (m *Mixer) BufferNext(n int) []byte {
	var buffer []byte
	for b := 0; b < n; b+=2 {
		// this is 16-bit big-endian; TODO: support different bits and platform byte order.
		buffer = append(buffer, m.mixNextHzBytes()...)
	}
	return buffer
}

func (m *Mixer) Teardown() {
	// nothing yet
}

/*
 *
 private */

func (m *Mixer) mixNextHz() smp16 {
	nowAtDur := time.Duration(m.nowAtHz) * time.Second / time.Duration(m.spec.Freq)
	m.nowAtHz++
	var Z []smp16
	for _, fire := range m.fire {
		if sourceAt := fire.At(m.freq, nowAtDur); sourceAt > 0 {
			m.Printf("play %s at %+v\n", fire.Source(), sourceAt)
			Z = append(Z, m.sourceAt(fire.Source(), sourceAt))
		}
	}
	sumZ := smp16(0)
	for _, z := range Z {
		sumZ += z
	}
	if cntZ := smp16(len(Z)); cntZ > 0 {
		return sumZ / cntZ
	} else {
		return smp16(m.spec.Silence)
	}
}

func (m *Mixer) mixNextHzBytes() []byte {
	// TODO: dynamically support modes other than 16-bit Big-Endian (which is coded below as b+=2 and the two buffer[..] assignments per chunk)
	byZ := make([]byte, 2)
	binary.BigEndian.PutUint16(byZ, uint16(m.mixNextHz()))
	return byZ
}

func (m *Mixer) sourceAt(source string, at Hz) smp16 {
	src := m.getSource(source)
	if src == nil {
		return 0x80
	}
	return src.At(at)
}

func (m *Mixer) setSpec(s sdl.AudioSpec) {
	m.spec = s
	m.freq = float64(s.Freq) // cache a float64 of this for future maths
}

func (m *Mixer) getSpec() *sdl.AudioSpec {
	return &m.spec
}

func (m *Mixer) prepareSource(source string) {
	// TODO: prepare this source and store in the map by its string
	if _, ok := m.source[source]; ok {
		// exists; take no action
	} else {
		// not exist
		m.source[source] = NewSource(source)
	}
}

func (m *Mixer) getSource(source string) *Source {
	if _, ok := m.source[source]; ok {
		return m.source[source]
	} else {
		return nil
	}
}

/*

// read audio file

var storedAudio []C.Uint16
var	(
	defaultSample = smp16(0xFFFF)
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
