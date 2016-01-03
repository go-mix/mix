/** Copyright 2015 Outright Mental, Inc. */
package atomix // is for sequence mixing

import (
	"github.com/stretchr/testify/assert"
	"github.com/veandco/go-sdl2/sdl"
	"testing"
)

//
// Tests
//

func TestNewMixer(t *testing.T) {
	assert.NotNil(t, Spec(&sdl.AudioSpec{
		Freq:     44100,
		Format:   sdl.AUDIO_U16,
		Channels: 2,
		Samples:  4096,
	}))
}

func TestNewMixerFailsWithoutProperAudioSpec(t *testing.T) {
	assert.Panics(t, func() {
		Spec(&sdl.AudioSpec{})
	})
}

//
// Components (to support Testing)
//
