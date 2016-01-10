// Mixer singleton orchestrates Sources and Fires
// Copyright 2015 Outright Mental, Inc.
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

func Test_Mixer_Initialize(t *testing.T) {
	// TODO: Test Mixer Initialize
}

func Test_Mixer_Debug(t *testing.T) {
	// TODO: Test Mixer Debug
}

func Test_Mixer_Debugf(t *testing.T) {
	// TODO: Test Mixer Debugf
}

func Test_Mixer_Start(t *testing.T) {
	// TODO: Test Mixer Start
}

func Test_Mixer_StartAt(t *testing.T) {
	// TODO: Test Mixer StartAt
}

func Test_Mixer_SetFire(t *testing.T) {
	// TODO: Test Mixer SetFire
}

func Test_Mixer_NextOutput(t *testing.T) {
	// TODO: Test Mixer NextOutput
}

func Test_Mixer_Teardown(t *testing.T) {
	// TODO: Test Mixer Teardown
}

func Test_Mixer_nextSample(t *testing.T) {
	// TODO: Test Mixer nextSample
}

func Test_Mixer_sourceAtTz(t *testing.T) {
	// TODO: Test Mixer sourceAtTz
}

func Test_Mixer_setSpec(t *testing.T) {
	// TODO: Test Mixer setSpec
}

func Test_Mixer_getSpec(t *testing.T) {
	// TODO: Test Mixer getSpec
}

func Test_Mixer_prepareSource(t *testing.T) {
	// TODO: Test Mixer prepareSource
}

func Test_Mixer_getSource(t *testing.T) {
	// TODO: Test Mixer getSource
}

func Test_Mixer_mix8(t *testing.T) {
	// TODO: Test Mixer mix8
}

func Test_Mixer_mix16(t *testing.T) {
	// TODO: Test Mixer mix16
}

func Test_Mixer_mix32(t *testing.T) {
	// TODO: Test Mixer mix32
}

func Test_mixByteU8(t *testing.T) {
	// TODO: Test mixByteU8
}

func Test_mixByteS8(t *testing.T) {
	// TODO: Test mixByteS8
}

func Test_mixBytesU16LSB(t *testing.T) {
	// TODO: Test mixBytesU16LSB
}

func Test_mixBytesU16MSB(t *testing.T) {
	// TODO: Test mixBytesU16MSB
}

func Test_mixBytesS16LSB(t *testing.T) {
	// TODO: Test mixBytesS16LSB
}

func Test_mixBytesS16MSB(t *testing.T) {
	// TODO: Test mixBytesS16MSB
}

func Test_mixBytesS32LSB(t *testing.T) {
	// TODO: Test mixBytesS32LSB
}

func Test_mixBytesS32MSB(t *testing.T) {
	// TODO: Test mixBytesS32MSB
}

func Test_mixBytesF32LSB(t *testing.T) {
	// TODO: Test mixBytesF32LSB
}

func Test_mixBytesF32MSB(t *testing.T) {
	// TODO: Test mixBytesF32MSB
}

func Test_mixUint8(t *testing.T) {
	// TODO: Test mixUint8
}

func Test_mixInt8(t *testing.T) {
	// TODO: Test mixInt8
}

func Test_mixUint16(t *testing.T) {
	// TODO: Test mixUint16
}

func Test_mixInt16(t *testing.T) {
	// TODO: Test mixInt16
}

func Test_mixUint32(t *testing.T) {
	// TODO: Test mixUint32
}

func Test_mixInt32(t *testing.T) {
	// TODO: Test mixInt32
}

func Test_mixFloat32(t *testing.T) {
	// TODO: Test mixFloat32
}

// TODO: test atomix.GetSpec()

// TODO: test atomix.Debug(true) and atomix.Debug(false)

// TODO: test atomix.Play("filename", time, duration, volume)

// TODO: test sources are queued and loaded properly

// TODO: test audio sources are mixed properly into buffer

// TODO: test different timing of ^

// TODO: test different audio format / bitrate / samples of ^

// TODO: test buffer properly reported to AudioCallback
