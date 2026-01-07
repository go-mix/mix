// Package bind is for modular binding of mix to audio interface
package bind

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-mix/mix/bind/opt"
	"github.com/go-mix/mix/bind/spec"
)

func TestAPI_Configure(t *testing.T) {
	// Test Configure doesn't panic with valid spec
	s := spec.AudioSpec{
		Freq:     44100,
		Format:   spec.AudioF32,
		Channels: 2,
	}
	assert.NotPanics(t, func() {
		Configure(s)
	})
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
	UseOutput(opt.OutputNull)
	assert.Equal(t, opt.OutputNull, useOutput)
}

func TestAPI_UseOutputString(t *testing.T) {
	UseOutputString("wav")
	assert.Equal(t, opt.OutputWAV, useOutput)
}

func TestAPI_UseOutputString_Fail(t *testing.T) {
	defer func() {
		msg := recover()
		assert.IsType(t, "", msg)
		assert.Equal(t, "No such Output: this-will-panic", msg)
	}()
	UseOutputString("this-will-panic")
}

func TestAPI_IsDirectOutput(t *testing.T) {
	// Test IsDirectOutput with WAV output
	UseOutput(opt.OutputWAV)
	assert.True(t, IsDirectOutput())

	// Test IsDirectOutput with null output
	UseOutput(opt.OutputNull)
	assert.False(t, IsDirectOutput())
}

func TestAPI_UseLoaderSOX(t *testing.T) {
	UseLoader(opt.InputSOX)
	assert.Equal(t, opt.InputSOX, useLoader)
}

func TestAPI_UseLoaderSOXString(t *testing.T) {
	UseLoaderString("sox")
	assert.Equal(t, opt.InputSOX, useLoader)
}

func TestAPI_Teardown(t *testing.T) {
	// Test teardown doesn't panic
	UseOutput(opt.OutputNull)
	assert.NotPanics(t, func() {
		Teardown()
	})

	UseOutput(opt.OutputWAV)
	assert.NotPanics(t, func() {
		Teardown()
	})
}
