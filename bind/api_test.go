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
	UseLoader(OptLoaderWAV)
	assert.Equal(t, OptLoaderWAV, useLoader)
}

func TestAPI_UseWAVString(t *testing.T) {
	UseLoaderString("wav")
	assert.Equal(t, OptLoaderWAV, useLoader)
}

func TestAPI_UseWAVString_Fail(t *testing.T) {
	defer func() {
		msg := recover()
		assert.IsType(t, "", msg)
		assert.Equal(t, "No such Loader: this-will-panic", msg)
	}()
	UseLoaderString("this-will-panic")
}

func TestAPI_UsePlayback(t *testing.T) {
	UsePlayback(OptPlaybackPortaudio)
	assert.Equal(t, OptPlaybackPortaudio, usePlayback)
	UsePlayback(OptPlaybackSDL)
	assert.Equal(t, OptPlaybackSDL, usePlayback)
}

func TestAPI_UsePlaybackString(t *testing.T) {
	UsePlaybackString("portaudio")
	assert.Equal(t, OptPlaybackPortaudio, usePlayback)
	UsePlaybackString("sdl")
	assert.Equal(t, OptPlaybackSDL, usePlayback)
}

func TestAPI_UsePlaybackString_Fail(t *testing.T) {
	defer func() {
		msg := recover()
		assert.IsType(t, "", msg)
		assert.Equal(t, "No such Playback: this-will-panic", msg)
	}()
	UsePlaybackString("this-will-panic")
}

func TestAPI_noErr(t *testing.T) {
	//TODO: Test
}

func TestAPI_sampleByteU8(t *testing.T) {
	//TODO: Test
}

func TestAPI_sampleByteS8(t *testing.T) {
	//TODO: Test
}

func TestAPI_sampleBytesU16LSB(t *testing.T) {
	//TODO: Test
}

func TestAPI_sampleBytesU16MSB(t *testing.T) {
	//TODO: Test
}

func TestAPI_sampleBytesS16LSB(t *testing.T) {
	//TODO: Test
}

func TestAPI_sampleBytesS16MSB(t *testing.T) {
	//TODO: Test
}

func TestAPI_sampleBytesS32LSB(t *testing.T) {
	//TODO: Test
}

func TestAPI_sampleBytesS32MSB(t *testing.T) {
	//TODO: Test
}

func TestAPI_sampleBytesF32LSB(t *testing.T) {
	//TODO: Test
}

func TestAPI_sampleBytesF32MSB(t *testing.T) {
	//TODO: Test
}

func TestAPI_sampleBytesF64LSB(t *testing.T) {
	//TODO: Test
}

func TestAPI_sampleBytesF64MSB(t *testing.T) {
	//TODO: Test
}
