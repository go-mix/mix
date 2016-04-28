// Package bind is for modular binding of mix to audio interface
package bind

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-mix/mix/bind/opt"
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

func TestAPI_noErr(t *testing.T) {
	//TODO: Test
}
