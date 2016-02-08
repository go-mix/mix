// Package atomix is a sequence-based Go-native audio mixer
package atomix

import (
	"github.com/outrightmental/go-atomix/bind"
	"math"
)

// NewSource from a "URL" (which is actually only a file path for now)
func NewSource(URL string) *Source {
	// TODO: implement true URL (for now, it's being used as a path)
	s := &Source{
		URL: URL,
	}
	s.state = sourceStateLoading
	s.load()
	s.state = sourceStateReady
	return s
}

// Source stores a series of Samples in Channels across Time, for audio playback.
type Source struct {
	URL    string
	sample [][]float64
	maxTz  Tz
	spec   *bind.AudioSpec
	state  sourceStateEnum
}

// SampleAt at a specific Tz, volume (0 to 1), and pan (-1 to +1)
func (s *Source) SampleAt(at Tz, volume float64, pan float64) (out []float64) {
	out = make([]float64, mixSpec.Channels)
	if at < s.maxTz {
		// if s.sample[at] != 0 {
		// 	mixDebugf("*Source[%v].SampleAt(%v): %v\n", s.URL, at, s.sample[at])
		// }
		if mixSpec.Channels == s.spec.Channels { // same # channels; easier maths
			for c := int(0); c < mixSpec.Channels; c++ {
				out[c] = mixVolume(float64(c), volume, pan) * s.sample[at][c]
			}
		} else { // need to map # source channels to # destination channels
			tc := float64(s.spec.Channels)
			for c := int(0); c < mixSpec.Channels; c++ {
				out[c] = mixVolume(float64(c), volume, pan) * s.sample[at][int(math.Floor(tc*float64(c)/mixChannels))]
			}
		}
	}
	return
}

// Length of the source audio in Tz
func (s *Source) Length() Tz {
	return s.maxTz
}

// Spec of the source audio
func (s *Source) Spec() *bind.AudioSpec {
	return s.spec
}

// Teardown the source audio and release its memory.
func (s *Source) Teardown() {
	s.sample = nil
}

/*
 *
 private */

func (s *Source) load() {
	s.sample, s.spec = bind.LoadWAV(s.URL)
	if s.spec == nil {
		// TODO: handle errors loading file
		mixDebugf("could not load WAV %s\n", s.URL)
	}
	s.maxTz = Tz(len(s.sample))
}

type sourceStateEnum uint

const (
	sourceStateLoading sourceStateEnum = 1
	sourceStateReady   sourceStateEnum = 2
	// it is assumed that all alive states are < SOURCE_FINISHED
	sourceStateFinished sourceStateEnum = 6
	sourceStateFailed   sourceStateEnum = 7
)
