// Package bind is for modular binding of atomix to audio interface
package bind

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAPI(t *testing.T) {
	// TODO
}

func TestAPI_UseWAV(t *testing.T) {
	UseWAV("native")
	assert.Equal(t, OptWAVGo, useWAV)
}

func TestAPI_UseWAV_Fail(t *testing.T) {
	defer func() {
		msg := recover()
		assert.IsType(t, "", msg)
		assert.Equal(t, "No such WAV: this-will-panic", msg)
	}()
	UseWAV("this-will-panic")
}

func TestAPI_UsePlayback(t *testing.T) {
	UsePlayback("portaudio")
	assert.Equal(t, OptPlaybackPortaudio, usePlayback)
	UsePlayback("sdl")
	assert.Equal(t, OptPlaybackSDL, usePlayback)
}

func TestAPI_UsePlayback_Fail(t *testing.T) {
	defer func() {
		msg := recover()
		assert.IsType(t, "", msg)
		assert.Equal(t, "No such Playback: this-will-panic", msg)
	}()
	UsePlayback("this-will-panic")
}
