// Package atomix is a sequence-based Go-native audio mixer
package atomix

import (
	"github.com/outrightmental/go-atomix/bind"
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

// SampleAt at a specific Tz (currently only mono, uses channel 0)
func (s *Source) SampleAt(at Tz) float64 {
	if at < s.maxTz {
		// if s.sample[at] != 0 {
		// 	Debugf("*Source[%v].SampleAt(%v): %v\n", s.URL, at, s.sample[at])
		// }
		return s.sample[at][0]
	}
	return 0
}

// Length of the source audio in Tz
func (s *Source) Length() Tz {
	return s.maxTz
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
