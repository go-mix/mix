// API exposes methods for use
package atomix // is for sequence mixing
// Copyright 2015 Outright Mental, Inc.

import (
	"github.com/stretchr/testify/assert"
	"github.com/veandco/go-sdl2/sdl"
	"time"
	"testing"
)

func TestAPI_Debug(t *testing.T) {
	// TODO: Test API Debug
}

func TestAPI_Configure(t *testing.T) {
	// TODO: Test API Configure
}

func TestAPI_Teardown(t *testing.T) {
	// TODO: Test API Teardown
}

func TestAPI_Spec(t *testing.T) {
	// TODO: Test API Spec
}

func TestAPI_SetFire(t *testing.T) {
	testApiSetup()
	fire := SetFire("lib/S16.wav", time.Duration(0), 0, 1.0, 0)
	assert.NotNil(t, fire)
}

func TestAPI_SetSoundsPath(t *testing.T) {
	// TODO: Test API SetSoundsPath
}

func TestAPI_Start(t *testing.T) {
	// TODO: Test API Start
}

func TestAPI_StartAt(t *testing.T) {
	// TODO: Test API StartAt
}

func TestAPI_GetStartTime(t *testing.T) {
	// TODO: Test API GetStartTime
}

func TestAPI_AudioCallback(t *testing.T) {
	// TODO: Test API AudioCallback
}

//
// Test Components
//

func testApiSetup() {
	Configure(sdl.AudioSpec{
		Freq:     44100,
		Format:   sdl.AUDIO_S16,
		Channels: 1,
		Samples:  4096,
	})
}
