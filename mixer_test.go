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

func Test_Mixer_Base(t *testing.T) {
	Configure(sdl.AudioSpec{
		Freq:     44100,
		Format:   sdl.AUDIO_U16,
		Channels: 2,
		Samples:  4096,
	})
	assert.NotNil(t, Spec())
}

func Test_Mixer_RequiresProperAudioSpec(t *testing.T) {
	assert.Panics(t, func() {
		Configure(sdl.AudioSpec{})
	})
}

// TODO: test atomix.GetSpec()

// TODO: test atomix.Debug(true) and atomix.Debug(false)

// TODO: test atomix.Play("filename", time, duration, volume)

// TODO: test sources are queued and loaded properly

// TODO: test audio sources are mixed properly into buffer

// TODO: test different timing of ^

// TODO: test different audio format / bitrate / samples of ^

// TODO: test buffer properly reported to AudioCallback

//
// Components (to support Testing)
//
