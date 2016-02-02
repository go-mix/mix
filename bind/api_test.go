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
	UseWAV(OptWAVGo)
	assert.Equal(t, OptWAVGo, useWAV)
}

func TestAPI_UsePlayback(t *testing.T) {
	UsePlayback(OptPlaybackPortaudio)
	assert.Equal(t, OptPlaybackPortaudio, usePlayback)
}
