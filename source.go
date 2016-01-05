/** Copyright 2015 Outright Mental, Inc. */
package atomix // is for sequence mixing

import (
	"encoding/binary"
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

func NewSource(URL string) *Source {
	// TODO: implement true URL (for now, it's being used as a path)
	s := &Source{
		URL: URL,
	}
	s.state = SOURCE_LOADING
	s.Load()
	s.state = SOURCE_READY
	return s
}

type Source struct {
	URL    string
	sample []float64
	spec   *sdl.AudioSpec
	state  SourceStateEnum
}

func (s *Source) SampleAt(at Tz) float64 {
	if at >= 0 && at < Tz(len(s.sample)) {
		return s.sample[at]
	}
	return 0
}

func (s *Source) State() SourceStateEnum {
	return s.state
}

func (s *Source) StateName() string {
	switch s.state {
	case SOURCE_LOADING:
		return "Loading"
	case SOURCE_READY:
		return "Ready"
	case SOURCE_FINISHED:
		return "Finished"
	case SOURCE_FAILED:
		return "Failed"
	}
	return ""
}

func (s *Source) Teardown() {
	s.sample = nil
}

/*
 *
 private */

func (s *Source) Load() {
	// TODO: support audio formats other than WAV
	data, spec := sdl.LoadWAV(s.URL, &sdl.AudioSpec{})
	if spec == nil || spec.Format == 0 {
		// TODO: handle errors loading file
		mixer().Debugf("could not load WAV %s", s.URL)
	}
	s.spec = spec
	switch s.spec.Format {
	case
		sdl.AUDIO_U8,
		sdl.AUDIO_S8:
		s.load8(data)
	case
		sdl.AUDIO_U16LSB,
		sdl.AUDIO_S16LSB,
		sdl.AUDIO_U16MSB,
		sdl.AUDIO_S16MSB:
		s.load16(data)
	case
		sdl.AUDIO_S32LSB,
		sdl.AUDIO_S32MSB,
		sdl.AUDIO_F32LSB:
		s.load32(data)
	default:
		mixer().Debugf("could not load WAV format %+v", s.spec.Format)
	}
}

func (s *Source) load8(data []byte) {
	// TODO: convert source Tz to the mixer output Tz
	for n := 0; n < len(data); n++ {
		switch s.spec.Format {
		case sdl.AUDIO_U8:
			s.sample = append(s.sample, sampleByteU8(data[n]))
		case sdl.AUDIO_S8:
			s.sample = append(s.sample, sampleByteS8(data[n]))
		}
	}
}

func (s *Source) load16(data []byte) {
	// TODO: convert source Tz to the mixer output Tz
	for n := 0; n < len(data); n += 2 {
		switch s.spec.Format {
		case sdl.AUDIO_U16LSB:
			s.sample = append(s.sample, sampleBytesU16LSB(data[n:n+2]))
		case sdl.AUDIO_S16LSB:
			s.sample = append(s.sample, sampleBytesS16LSB(data[n:n+2]))
		case sdl.AUDIO_U16MSB:
			s.sample = append(s.sample, sampleBytesU16MSB(data[n:n+2]))
		case sdl.AUDIO_S16MSB:
			s.sample = append(s.sample, sampleBytesS16MSB(data[n:n+2]))
		}
	}
}

func (s *Source) load32(data []byte) {
	// TODO: convert source Tz to the mixer output Tz
	for n := 0; n < len(data); n += 4 {
		switch s.spec.Format {
		case sdl.AUDIO_S32LSB:
			s.sample = append(s.sample, sampleBytesS32LSB(data[n:n+4]))
		case sdl.AUDIO_S32MSB:
			s.sample = append(s.sample, sampleBytesS32MSB(data[n:n+4]))
		case sdl.AUDIO_F32LSB:
			s.sample = append(s.sample, sampleBytesF32LSB(data[n:n+4]))
		case sdl.AUDIO_F32MSB:
			s.sample = append(s.sample, sampleBytesF32MSB(data[n:n+4]))
		}
	}
}

func sampleByteU8(sample byte) float64 {
	return float64(int8(sample)/0x7F - 1)
}

func sampleByteS8(sample byte) float64 {
	return float64(int8(sample) / 0x7F)
}

func sampleBytesU16LSB(sample []byte) float64 {
	return float64(binary.LittleEndian.Uint16(sample)/0x8000 - 1)
}

func sampleBytesU16MSB(sample []byte) float64 {
	return float64(binary.BigEndian.Uint16(sample)/0x8000 - 1)
}

func sampleBytesS16LSB(sample []byte) float64 {
	return float64(int16(binary.LittleEndian.Uint16(sample)) / 0x7FFF)
}

func sampleBytesS16MSB(sample []byte) float64 {
	return float64(int16(binary.BigEndian.Uint16(sample)) / 0x7FFF)
}

func sampleBytesS32LSB(sample []byte) float64 {
	return float64(int32(binary.LittleEndian.Uint32(sample)) / 0x7FFFFFFF)
}

func sampleBytesS32MSB(sample []byte) float64 {
	return float64(int32(binary.BigEndian.Uint32(sample)) / 0x7FFFFFFF)
}

func sampleBytesF32LSB(sample []byte) float64 {
	return float64(math.Float32frombits(binary.LittleEndian.Uint32(sample)))
}

func sampleBytesF32MSB(sample []byte) float64 {
	return float64(math.Float32frombits(binary.BigEndian.Uint32(sample)))
}

type SourceStateEnum uint

const (
	SOURCE_LOADING SourceStateEnum = 1
	SOURCE_READY   SourceStateEnum = 2
	// it can be assumed that all alive states are < SOURCE_FINISHED
	SOURCE_FINISHED SourceStateEnum = 6
	SOURCE_FAILED   SourceStateEnum = 7
)
