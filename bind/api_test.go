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

func TestAPI_UseOutput(t *testing.T) {
	UseOutput(OptOutputPortaudio)
	assert.Equal(t, OptOutputPortaudio, useOutput)
	UseOutput(OptOutputSDL)
	assert.Equal(t, OptOutputSDL, useOutput)
}

func TestAPI_UseOutputString(t *testing.T) {
	UseOutputString("portaudio")
	assert.Equal(t, OptOutputPortaudio, useOutput)
	UseOutputString("sdl")
	assert.Equal(t, OptOutputSDL, useOutput)
}

func TestAPI_UseOutputString_Fail(t *testing.T) {
	defer func() {
		msg := recover()
		assert.IsType(t, "", msg)
		assert.Equal(t, "No such Output: this-will-panic", msg)
	}()
	UseOutputString("this-will-panic")
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
