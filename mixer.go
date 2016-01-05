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

type Hz    uint64

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

func (m *Mixer) NextOutputBytes() []byte {
	return m.mixNextHzBytes()
}

func (m *Mixer) Teardown() {
	// nothing yet
}

/*
 *
 private */

func (m *Mixer) mixNextHz() uint16 {
	nowAtDur := time.Duration(m.nowAtHz) * time.Second / time.Duration(m.spec.Freq)
	m.nowAtHz++
	var mixThisHz []uint16
	for _, fire := range m.fire {
		if fireHz := fire.NextHzAt(nowAtDur); fireHz > 0 {
			mixThisHz = append(mixThisHz, m.sourceAtHz(fire.source, fireHz))
		}
	}
	mixSum := uint16(0)
	for _, sample := range mixThisHz {
		mixSum += sample
	}
	if mixCount := uint16(len(mixThisHz)); mixCount > 0 {
		return mixSum / mixCount
	} else {
		return uint16(m.spec.Silence)
	}
}

func (m *Mixer) mixNextHzBytes() []byte {
	b := make([]byte, 2)
	// TODO: dynamically support modes other than 16-bit Big-Endian
	binary.BigEndian.PutUint16(b, m.mixNextHz())
	return b
}

func (m *Mixer) sourceAtHz(src string, srcHz Hz) uint16 {
	s := m.getSource(src)
	if s == nil {
		return 0x80
	}
	return s.SampleAt(srcHz)
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
