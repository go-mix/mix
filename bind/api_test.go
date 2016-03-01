// Package bind is for modular binding of ontomix to audio interface
package bind

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gopkg.in/ontomix.v0/bind/opt"
)

func TestAPI(t *testing.T) {
	// TODO
}

func TestAPI_UseWAV(t *testing.T) {
	UseLoader(opt.InputWAV)
	assert.Equal(t, opt.InputWAV, useLoader)
}

func TestAPI_UseWAVString(t *testing.T) {
	UseLoaderString("wav")
	assert.Equal(t, opt.InputWAV, useLoader)
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
	UseOutput(opt.OutputPortAudio)
	assert.Equal(t, opt.OutputPortAudio, useOutput)
	UseOutput(opt.OutputSDL)
	assert.Equal(t, opt.OutputSDL, useOutput)
}

func TestAPI_UseOutputString(t *testing.T) {
	UseOutputString("portaudio")
	assert.Equal(t, opt.OutputPortAudio, useOutput)
	UseOutputString("sdl")
	assert.Equal(t, opt.OutputSDL, useOutput)
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
