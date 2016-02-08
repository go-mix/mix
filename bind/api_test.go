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
	UseLoader("wav")
	assert.Equal(t, OptLoaderWAV, useLoader)
}

func TestAPI_UseWAV_Fail(t *testing.T) {
	defer func() {
		msg := recover()
		assert.IsType(t, "", msg)
		assert.Equal(t, "No such Loader: this-will-panic", msg)
	}()
	UseLoader("this-will-panic")
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
