/** Copyright 2015 Outright Mental, Inc. */
package atomix // is for sequence mixing

import (
	"fmt"
	"encoding/binary"
	"github.com/veandco/go-sdl2/sdl"
)

func NewSource(target string) *Source {
	s := &Source{
		target: target,
	}
	s.state = STATE_LOADING
	s.loadFile()
	s.state = STATE_READY
	return s
}

type Source struct {
	target         string
	sample         []uint16
	spec           *sdl.AudioSpec
	state          SourceStateEnum
}

func (s *Source) SampleAt(at Hz) uint16 {
	if at >= 0 && at < Hz(len(s.sample)) {
		return s.sample[at]
	}
	return 0
}

func (s *Source) State() SourceStateEnum {
	return s.state
}

func (s *Source) StateName() string {
	switch s.state {
	case STATE_LOADING:
		return "Loading"
	case STATE_READY:
		return "Ready"
	case STATE_FINISHED:
		return "Finished"
	case STATE_FAILED:
		return "Failed"
	}
	return ""
}

func (s *Source) Teardown() {
	// TODO: free memory of s.sample
}

/*
 *
 private */

func (s *Source) loadFile() {
	// TODO: handle errors loading file (panic)
	data, spec := sdl.LoadWAV(s.target, &sdl.AudioSpec{})
	if spec == nil || spec.Format==0 {
		panic(fmt.Sprintf("could not load WAV %s", s.target))
	}
	for n := 0; n < len(data); n += 2 {
		s.sample = append(s.sample, binary.BigEndian.Uint16(data[n:n+2]))
	}
	s.spec = spec
}

type SourceStateEnum uint

const (
	STATE_LOADING  SourceStateEnum = 1
	STATE_READY    SourceStateEnum = 2
	// it can be assumed that all alive states are < STATE_FINISHED
	STATE_FINISHED SourceStateEnum = 6
	STATE_FAILED   SourceStateEnum = 7
)
