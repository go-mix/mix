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
	sample         []smp16
	state          SourceStateEnum
}

func (s *Source) At(at Hz) smp16 {
	// TODO: dynamically support modes other than 16-bit (which is coded below as 2 * at and the two value slice)
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
	s.store(data)
	mixer().Printf("Stored %d bytes for '%s'\n", len(s.sample), s.target)
}

func (s *Source) store(d []byte) {
	s.sample = append(s.sample, smp16(binary.BigEndian.Uint16(d)))
}

type SourceStateEnum uint

const (
	STATE_LOADING  SourceStateEnum = 1
	STATE_READY    SourceStateEnum = 2
	// it can be assumed that all alive states are < STATE_FINISHED
	STATE_FINISHED SourceStateEnum = 6
	STATE_FAILED   SourceStateEnum = 7
)
