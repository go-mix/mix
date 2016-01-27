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
	mixMutex       *sync.Mutex = &sync.Mutex{}
	mixStartAtTime time.Time
	mixNowTz       Tz
	mixTzDur       time.Duration
	// TODO: implement mixFreq float64
	mixSource       map[string]*Source
	mixSourcePrefix string
	mixFires        []*Fire
	mixSpec         sdl.AudioSpec
	isDebug         bool
)

type Tz uint64

func mixDebug(isOn bool) {
	isDebug = isOn
}

func mixDebugf(format string, args ...interface{}) {
	if isDebug {
		fmt.Printf(format, args...)
	}
}

func mixStartAt(t time.Time) {
	mixStartAtTime = t
}

func mixGetStartTime() time.Time {
	return mixStartAtTime
}

func mixSourceLength(source string) Tz {
	s := mixGetSource(source)
	if s == nil {
		return 0
	}
	return s.Length()
}

func mixSetFire(source string, begin time.Duration, sustain time.Duration, volume float64, pan float64) *Fire {
	mixPrepareSource(mixSourcePrefix + source)
	beginTz := Tz(begin.Nanoseconds() / mixTzDur.Nanoseconds())
	var endTz Tz
	if sustain != 0 {
		endTz = beginTz + Tz(sustain.Nanoseconds()/mixTzDur.Nanoseconds())
	}
	fire := NewFire(mixSourcePrefix+source, beginTz, endTz, volume, pan)
	mixFires = append(mixFires, fire)
	return fire
}

func mixSetSoundsPath(prefix string) {
	mixSourcePrefix = prefix
}

func mixNextOutput(byteSize int) []byte {
	mixCleanup()
	switch mixSpec.Format {
	case
		sdl.AUDIO_U8,
		sdl.AUDIO_S8:
		return mix8(byteSize)
	case
		sdl.AUDIO_U16LSB,
		sdl.AUDIO_S16LSB,
		sdl.AUDIO_U16MSB,
		sdl.AUDIO_S16MSB:
		return mix16(byteSize)
	case
		sdl.AUDIO_S32LSB,
		sdl.AUDIO_S32MSB,
		sdl.AUDIO_F32LSB:
		return mix32(byteSize)
	}
	return nil
}

func mixTeardown() {
	// nothing yet
}

/*
 *
 private */

func mixNextSample() float64 {
	sample := float64(0)
	// TODO: #FIXME need a more efficient method of iterating active fires; range fires hogs CPU with >100 fires
	// TODO: #FIXME ^ really this is a serious processor bottleneck. Find a method to avoid iterating over all these inactive fires every sample!
	for _, fire := range mixFires {
		// mixer().Debugf("see me try to fire? %v", fire.Source, fire.BeginTz)
		if fireTz := fire.At(mixNowTz); fireTz > 0 {
			sample += mixSourceAt(fire.Source, fireTz)
		}
	}
	// if sample != 0 {
	// 	Debugf("*Mixer.nextSample at %+v: %+v\n", nowTz, sample)
	// }
	mixNowTz++
	return mixLogarithmicRangeCompression(sample)
}

func mixSourceAt(src string, at Tz) float64 {
	s := mixGetSource(src)
	if s == nil {
		return 0
	}
	// if at != 0 {
	// 	Debugf("About to source.SampleAt %v in %v\n", at, s.URL)
	// }
	return s.SampleAt(at)
}

func mixSetSpec(s sdl.AudioSpec) {
	mixSpec = s
	// TODO: implement mixFreq = float64(s.Freq) // cache a float64 of this for future maths
	mixTzDur = time.Second / time.Duration(s.Freq)
}

func mixGetSpec() *sdl.AudioSpec {
	return &mixSpec
}

func mixPrepareSource(source string) {
	mixMutex.Lock()
	defer mixMutex.Unlock()
	if _, exists := mixSource[source]; !exists {
		mixSource[source] = NewSource(source)
	}
}

func mixGetSource(source string) *Source {
	mixMutex.Lock()
	defer mixMutex.Unlock()
	if _, ok := mixSource[source]; ok {
		return mixSource[source]
	} else {
		return nil
	}
}

func mixCleanup() {
	for i, fire := range mixFires {
		if !fire.IsAlive() {
			fire.Teardown()
			mixFires = append(mixFires[:i], mixFires[i+1:]...)
		}
	}
}

func mix8(byteSize int) (out []byte) {
	for n := 0; n < byteSize; n++ {
		switch mixSpec.Format {
		case sdl.AUDIO_U8:
			out = append(out, mixByteU8(mixNextSample()))
		case sdl.AUDIO_S8:
			out = append(out, mixByteS8(mixNextSample()))
		}
	}
	return
}

func mix16(byteSize int) (out []byte) {
	for n := 0; n < byteSize; n += 2 {
		switch mixSpec.Format {
		case sdl.AUDIO_U16LSB:
			out = append(out, mixBytesU16LSB(mixNextSample())...)
		case sdl.AUDIO_S16LSB:
			out = append(out, mixBytesS16LSB(mixNextSample())...)
		case sdl.AUDIO_U16MSB:
			out = append(out, mixBytesU16MSB(mixNextSample())...)
		case sdl.AUDIO_S16MSB:
			out = append(out, mixBytesS16MSB(mixNextSample())...)
		}
	}
	return
}

func mix32(byteSize int) (out []byte) {
	for n := 0; n < byteSize; n += 4 {
		switch mixSpec.Format {
		case sdl.AUDIO_S32LSB:
			out = append(out, mixBytesS32LSB(mixNextSample())...)
		case sdl.AUDIO_S32MSB:
			out = append(out, mixBytesS32MSB(mixNextSample())...)
		case sdl.AUDIO_F32LSB:
			out = append(out, mixBytesF32LSB(mixNextSample())...)
		case sdl.AUDIO_F32MSB:
			out = append(out, mixBytesF32MSB(mixNextSample())...)
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

//func mixUint32(sample float64) uint32 {
//	return uint32(0x80000000 * (sample + 1))
//}

func mixInt32(sample float64) int32 {
	return int32(0x80000000 * sample)
}

//func mixFloat32(sample float64) float32 {
//	return float32(sample)
//}

func mixLogarithmicRangeCompression(i float64) float64 {
	if i < -1 {
		return -math.Log(-i-0.85)/14 - 0.75
	} else if i > 1 {
		return math.Log(i-0.85)/14 + 0.75
	} else {
		return i / 1.61803398875
	}
}

func init() {
	mixSource = make(map[string]*Source, 0)
	mixStartAtTime = time.Now().Add(0xFFFF * time.Hour) // this gets reset by Start() or StartAt()
}
