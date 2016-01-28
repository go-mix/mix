// Package atomix is a sequence-based Go-native audio mixer
package atomix

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
	sample [][]float64
	maxTz  Tz
	spec   *sdl.AudioSpec
	state  SourceStateEnum
}

func (s *Source) SampleAt(at Tz) float64 {
	if at < s.maxTz {
		// if s.sample[at] != 0 {
		// 	Debugf("*Source[%v].SampleAt(%v): %v\n", s.URL, at, s.sample[at])
		// }
		return s.sample[at][0]
	} else {
		return 0
	}
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

func (s *Source) Length() Tz {
	return s.maxTz
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
		mixDebugf("could not load WAV %s", s.URL)
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
		mixDebugf("could not load WAV format %+v", s.spec.Format)
	}
	s.maxTz = Tz(len(s.sample))
}

func (s *Source) load8(data []byte) {
	channels := int(s.spec.Channels)
	// TODO: convert source Hz; store at the mixer output Hz
	for n := 0; n < len(data); n++ {
		sample := make([]float64, channels)
		for c := 0; c < channels; c++ {
			switch s.spec.Format {
			case sdl.AUDIO_U8:
				sample[c] = sampleByteU8(data[n])
			case sdl.AUDIO_S8:
				sample[c] = sampleByteS8(data[n])
			}
		}
		// TODO: instead of append(..), make([][]float64,length) ahead of time!
		s.sample = append(s.sample, sample)
	}
	mixDebugf("*Source[%s].load8(...) length %d channels %d\n", s.URL, len(s.sample), s.spec.Channels)
}

func (s *Source) load16(data []byte) {
	channels := int(s.spec.Channels)
	// TODO: convert source Hz; store at the mixer output Hz
	for n := 0; n < len(data); n += 2 {
		sample := make([]float64, channels)
		for c := 0; c < channels; c++ {
			b := n + c*2
			switch s.spec.Format {
			case sdl.AUDIO_U16LSB:
				sample[c] = sampleBytesU16LSB(data[b : b+2])
			case sdl.AUDIO_S16LSB:
				sample[c] = sampleBytesS16LSB(data[b : b+2])
			case sdl.AUDIO_U16MSB:
				sample[c] = sampleBytesU16MSB(data[b : b+2])
			case sdl.AUDIO_S16MSB:
				sample[c] = sampleBytesS16MSB(data[b : b+2])
			}
		}
		// TODO: instead of append(..), make([][]float64,length) ahead of time!
		s.sample = append(s.sample, sample)
	}
	mixDebugf("*Source[%s].load16(...) length %d channels %d\n", s.URL, len(s.sample), s.spec.Channels)
}

func (s *Source) load32(data []byte) {
	channels := int(s.spec.Channels)
	// TODO: convert source Hz; store at the mixer output Hz
	for n := 0; n < len(data); n += channels * 4 {
		sample := make([]float64, channels)
		for c := 0; c < channels; c++ {
			b := n + c*4
			switch s.spec.Format {
			case sdl.AUDIO_S32LSB:
				sample[c] = sampleBytesS32LSB(data[b : b+4])
			case sdl.AUDIO_S32MSB:
				sample[c] = sampleBytesS32MSB(data[b : b+4])
			case sdl.AUDIO_F32LSB:
				sample[c] = sampleBytesF32LSB(data[b : b+4])
			case sdl.AUDIO_F32MSB:
				sample[c] = sampleBytesF32MSB(data[b : b+4])
			}
		}
		// TODO: instead of append(..), make([][]float64,length) ahead of time!
		s.sample = append(s.sample, sample)
	}
	mixDebugf("*Source[%s].load32(...) length %d channels %d\n", s.URL, len(s.sample), s.spec.Channels)
}

func sampleByteU8(sample byte) float64 {
	return float64(int8(sample))/float64(0x7F) - float64(1)
}

func sampleByteS8(sample byte) float64 {
	return float64(int8(sample)) / float64(0x7F)
}

func sampleBytesU16LSB(sample []byte) float64 {
	return float64(binary.LittleEndian.Uint16(sample))/float64(0x8000) - float64(1)
}

func sampleBytesU16MSB(sample []byte) float64 {
	return float64(binary.BigEndian.Uint16(sample))/float64(0x8000) - float64(1)
}

func sampleBytesS16LSB(sample []byte) float64 {
	return float64(int16(binary.LittleEndian.Uint16(sample))) / float64(0x7FFF)
}

func sampleBytesS16MSB(sample []byte) float64 {
	return float64(int16(binary.BigEndian.Uint16(sample))) / float64(0x7FFF)
}

func sampleBytesS32LSB(sample []byte) float64 {
	return float64(int32(binary.LittleEndian.Uint32(sample))) / float64(0x7FFFFFFF)
}

func sampleBytesS32MSB(sample []byte) float64 {
	return float64(int32(binary.BigEndian.Uint32(sample))) / float64(0x7FFFFFFF)
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
	// it is assumed that all alive states are < SOURCE_FINISHED
	SOURCE_FINISHED SourceStateEnum = 6
	SOURCE_FAILED   SourceStateEnum = 7
)
