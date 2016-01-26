// Mixer singleton orchestrates Sources and Fires
package atomix // is for sequence mixing
// Copyright 2015 Outright Mental, Inc.

/*
#include <stdio.h>
#include <stdint.h>
typedef unsigned char Uint8;
void AudioCallback(void *userdata, Uint8 *stream, int len);
*/
import "C"
import (
	"encoding/binary"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"math"
	"sync"
	"time"
)

var (
	defaultAudio = C.Uint8(0)
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

type Tz uint64

type Mixer struct {
	startAtTime  time.Time
	nowTz        Tz
	tzDur        time.Duration
	freq         float64
	source       map[string]*Source
	sourcePrefix string
	fires        []*Fire
	spec         sdl.AudioSpec
	isDebug      bool
}

func (m *Mixer) Initialize() {
	mutex := &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()
	m.source = make(map[string]*Source, 0)
	m.startAtTime = time.Now().Add(0xFFFF * time.Hour) // this gets reset by Start() or StartAt()
}

func (m *Mixer) Debug(isOn bool) {
	m.isDebug = isOn
}

func (m *Mixer) Debugf(format string, args ...interface{}) {
	if m.isDebug {
		fmt.Printf(format, args...)
	}
}

func (m *Mixer) StartAt(t time.Time) {
	m.startAtTime = t
}

func (m *Mixer) GetStartTime() time.Time {
	return m.startAtTime
}

func (m *Mixer) SourceLength(source string) Tz {
	s := m.getSource(source)
	if s == nil {
		return 0
	}
	return s.Length()
}

func (m *Mixer) SetFire(source string, begin time.Duration, sustain time.Duration, volume float64, pan float64) *Fire {
	m.prepareSource(m.sourcePrefix + source)
	beginTz := Tz(begin.Nanoseconds() / m.tzDur.Nanoseconds())
	var endTz Tz
	if sustain != 0 {
		endTz = beginTz + Tz(sustain.Nanoseconds()/m.tzDur.Nanoseconds())
	}
	fire := NewFire(m.sourcePrefix + source, beginTz, endTz, volume, pan)
	m.fires = append(m.fires, fire)
	return fire
}

func (m *Mixer) SetSoundsPath(prefix string) {
	m.sourcePrefix = prefix
}

func (m *Mixer) NextOutput(byteSize int) []byte {
	m.cleanup()
	switch m.spec.Format {
	case
		sdl.AUDIO_U8,
		sdl.AUDIO_S8:
		return m.mix8(byteSize)
	case
		sdl.AUDIO_U16LSB,
		sdl.AUDIO_S16LSB,
		sdl.AUDIO_U16MSB,
		sdl.AUDIO_S16MSB:
		return m.mix16(byteSize)
	case
		sdl.AUDIO_S32LSB,
		sdl.AUDIO_S32MSB,
		sdl.AUDIO_F32LSB:
		return m.mix32(byteSize)
	}
	return nil
}

func (m *Mixer) Teardown() {
	// nothing yet
}

/*
 *
 private */

func (m *Mixer) nextSample() float64 {
	sample := float64(0)
	// TODO: #FIXME need a more efficient method of iterating active fires; range m.fires hogs CPU with >100 fires
	// TODO: #FIXME ^ really this is a serious processor bottleneck. Find a method to avoid iterating over all these inactive fires every sample!
	for _, fire := range m.fires {
		// mixer().Debugf("see me try to fire? %v", fire.Source, fire.BeginTz)
		if fireTz := fire.At(m.nowTz); fireTz > 0 {
			sample += m.sourceAt(fire.Source, fireTz)
		}
	}
	// if sample != 0 {
	// 	m.Debugf("*Mixer.nextSample at %+v: %+v\n", m.nowTz, sample)
	// }
	m.nowTz++
	return mixLogarithmicRangeCompression(sample)
}

func (m *Mixer) sourceAt(src string, at Tz) float64 {
	s := m.getSource(src)
	if s == nil {
		return 0
	}
	// if at != 0 {
	// 	m.Debugf("About to source.SampleAt %v in %v\n", at, s.URL)
	// }
	return s.SampleAt(at)
}

func (m *Mixer) setSpec(s sdl.AudioSpec) {
	m.spec = s
	m.freq = float64(s.Freq) // cache a float64 of this for future maths
	m.tzDur = time.Second / time.Duration(s.Freq)
}

func (m *Mixer) getSpec() *sdl.AudioSpec {
	return &m.spec
}

func (m *Mixer) prepareSource(source string) {
	mutex := &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()
	if _, exists := m.source[source]; !exists {
		m.source[source] = NewSource(source)
	}
}

func (m *Mixer) getSource(source string) *Source {
	mutex := &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()
	if _, ok := m.source[source]; ok {
		return m.source[source]
	} else {
		return nil
	}
}

func (m *Mixer) cleanup() {
	for i, fire := range m.fires {
		if !fire.IsAlive() {
			fire.Teardown()
			m.fires = append(m.fires[:i], m.fires[i+1:]...)
		}
	}
}

func (m *Mixer) mix8(byteSize int) (out []byte) {
	for n := 0; n < byteSize; n++ {
		switch m.spec.Format {
		case sdl.AUDIO_U8:
			out = append(out, mixByteU8(m.nextSample()))
		case sdl.AUDIO_S8:
			out = append(out, mixByteS8(m.nextSample()))
		}
	}
	return
}

func (m *Mixer) mix16(byteSize int) (out []byte) {
	for n := 0; n < byteSize; n += 2 {
		switch m.spec.Format {
		case sdl.AUDIO_U16LSB:
			out = append(out, mixBytesU16LSB(m.nextSample())...)
		case sdl.AUDIO_S16LSB:
			out = append(out, mixBytesS16LSB(m.nextSample())...)
		case sdl.AUDIO_U16MSB:
			out = append(out, mixBytesU16MSB(m.nextSample())...)
		case sdl.AUDIO_S16MSB:
			out = append(out, mixBytesS16MSB(m.nextSample())...)
		}
	}
	return
}

func (m *Mixer) mix32(byteSize int) (out []byte) {
	for n := 0; n < byteSize; n += 4 {
		switch m.spec.Format {
		case sdl.AUDIO_S32LSB:
			out = append(out, mixBytesS32LSB(m.nextSample())...)
		case sdl.AUDIO_S32MSB:
			out = append(out, mixBytesS32MSB(m.nextSample())...)
		case sdl.AUDIO_F32LSB:
			out = append(out, mixBytesF32LSB(m.nextSample())...)
		case sdl.AUDIO_F32MSB:
			out = append(out, mixBytesF32MSB(m.nextSample())...)
		}
	}
	return
}

func mixByteU8(sample float64) byte {
	return byte(mixUint8(sample))
}

func mixByteS8(sample float64) byte {
	return byte(mixInt8(sample))
}

func mixBytesU16LSB(sample float64) (out []byte) {
	out = make([]byte, 2)
	binary.LittleEndian.PutUint16(out, mixUint16(sample))
	return
}

func mixBytesU16MSB(sample float64) (out []byte) {
	out = make([]byte, 2)
	binary.BigEndian.PutUint16(out, mixUint16(sample))
	return
}

func mixBytesS16LSB(sample float64) (out []byte) {
	out = make([]byte, 2)
	binary.LittleEndian.PutUint16(out, uint16(mixInt16(sample)))
	return
}

func mixBytesS16MSB(sample float64) (out []byte) {
	out = make([]byte, 2)
	binary.BigEndian.PutUint16(out, uint16(mixInt16(sample)))
	return
}

func mixBytesS32LSB(sample float64) (out []byte) {
	out = make([]byte, 4)
	binary.LittleEndian.PutUint32(out, uint32(mixInt32(sample)))
	return
}

func mixBytesS32MSB(sample float64) (out []byte) {
	out = make([]byte, 4)
	binary.BigEndian.PutUint32(out, uint32(mixInt32(sample)))
	return
}

func mixBytesF32LSB(sample float64) (out []byte) {
	out = make([]byte, 4)
	binary.LittleEndian.PutUint32(out, math.Float32bits(float32(sample)))
	return
}

func mixBytesF32MSB(sample float64) (out []byte) {
	out = make([]byte, 4)
	binary.BigEndian.PutUint32(out, math.Float32bits(float32(sample)))
	return
}

func mixUint8(sample float64) uint8 {
	return uint8(0x80 * (sample + 1))
}

func mixInt8(sample float64) int8 {
	return int8(0x80 * sample)
}

func mixUint16(sample float64) uint16 {
	return uint16(0x8000 * (sample + 1))
}

func mixInt16(sample float64) int16 {
	return int16(0x8000 * sample)
}

func mixUint32(sample float64) uint32 {
	return uint32(0x80000000 * (sample + 1))
}

func mixInt32(sample float64) int32 {
	return int32(0x80000000 * sample)
}

func mixFloat32(sample float64) float32 {
	return float32(sample)
}

func mixLogarithmicRangeCompression(i float64) float64 {
	if i < -1 {
		return -math.Log(-i - 0.85) / 14 - 0.75
	} else if i > 1 {
		return math.Log(i - 0.85) / 14 + 0.75
	} else {
		return i / 1.61803398875
	}
}
