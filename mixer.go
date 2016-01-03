/** Copyright 2015 Outright Mental, Inc. */
package atomix
/* in-buffer mixing when timing is known in advance (e.g. music) */

import (
  // log "github.com/Sirupsen/logrus"
	"github.com/veandco/go-sdl2/sdl"
)

type Mixer struct {
	/* private */
	spec *sdl.AudioSpec
}

// New creates a new Mixer
func New(spec sdl.AudioSpec) *Mixer {
  if spec.Freq == 0 {
    panic("Must specify Frequency")
  } else if spec.Format == 0 {
    panic("Must specify Format")
  } else if spec.Channels == 0 {
    panic("Must specify Channels")
  } else if spec.Samples == 0 {
    panic("Must specify Samples")
  }
	return &Mixer{
		spec: &spec,
	}
}

func (m *Mixer) GetSpec() *sdl.AudioSpec {
	return m.spec
}
