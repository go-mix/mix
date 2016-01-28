// Package atomix is a sequence-based Go-native audio mixer
package atomix

import (
	"github.com/stretchr/testify/assert"
	"github.com/veandco/go-sdl2/sdl"
	"testing"
)

//
// Tests
//

func TestMixer_Base(t *testing.T) {
	Configure(sdl.AudioSpec{
		Freq:     44100,
		Format:   sdl.AUDIO_U16,
		Channels: 2,
		Samples:  4096,
	})
	assert.NotNil(t, Spec())
}

func TestMixer_RequiresProperAudioSpec(t *testing.T) {
	assert.Panics(t, func() {
		Configure(sdl.AudioSpec{})
	})
}

func TestMixer_Initialize(t *testing.T) {
	// TODO: Test Mixer Initialize
}

func TestMixer_Debug(t *testing.T) {
	// TODO: Test Mixer Debug
}

func TestMixer_Debugf(t *testing.T) {
	// TODO: Test Mixer Debugf
}

func TestMixer_Start(t *testing.T) {
	// TODO: Test Mixer Start
}

func TestMixer_StartAt(t *testing.T) {
	// TODO: Test Mixer StartAt
}

func TestMixer_GetStartTime(t *testing.T) {
	// TODO: Test Mixer GetStartTime
}

func TestMixer_SetFire(t *testing.T) {
	// TODO: Test Mixer SetFire
}

func TestMixer_SetSoundsPath(t *testing.T) {
	// TODO: Test Mixer SetSoundsPath
}

func TestMixer_NextOutput(t *testing.T) {
	// TODO: Test Mixer NextOutput
}

func TestMixer_SourceLength(t *testing.T) {
	// TODO: Test Mixer SourceLength
}

func TestMixer_Teardown(t *testing.T) {
	// TODO: Test Mixer Teardown
}

func TestMixer_nextSample(t *testing.T) {
	// TODO: Test Mixer nextSample
}

func TestMixer_sourceAtTz(t *testing.T) {
	// TODO: Test Mixer sourceAt
}

func TestMixer_setSpec(t *testing.T) {
	// TODO: Test Mixer setSpec
}

func TestMixer_getSpec(t *testing.T) {
	// TODO: Test Mixer getSpec
}

func TestMixer_prepareSource(t *testing.T) {
	// TODO: Test Mixer prepareSource
}

func TestMixer_getSource(t *testing.T) {
	// TODO: Test Mixer getSource
}

func TestMixer_mix8(t *testing.T) {
	// TODO: Test Mixer mix8
}

func TestMixer_mix16(t *testing.T) {
	// TODO: Test Mixer mix16
}

func TestMixer_mix32(t *testing.T) {
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
