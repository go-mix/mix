/** Copyright 2015 Outright Mental, Inc. */
package atomix
/* in-buffer mixing when timing is known in advance (e.g. music) */

import (
	"github.com/stretchr/testify/assert"
	"github.com/veandco/go-sdl2/sdl"
	"testing"
)

//
// Tests
//

func TestNewMixer(t *testing.T) {
	mixer := New(sdl.AudioSpec{
		Freq:     44100,
		Format:   sdl.AUDIO_U16,
		Channels: 2,
		Samples:  4096,
	})
	assert.NotNil(t, mixer.GetSpec())
}

func TestNewMixerFailsWithoutProperAudioSpec(t *testing.T) {
	assert.Panics(t, func() {
		New(sdl.AudioSpec{})
	})
}

//
// Components (to support Testing)
//
