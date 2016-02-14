// Package atomix is a sequence-based Go-native audio mixer
package atomix

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/outrightmental/go-atomix/bind"
)

//
// Tests
//

func TestMixer_Base(t *testing.T) {
	Configure(bind.AudioSpec{
		Freq:     44100,
		Format:   bind.AudioU16,
		Channels: 2,
	})
	assert.NotNil(t, Spec())
}

func TestMixer_RequiresProperAudioSpec(t *testing.T) {
	assert.Panics(t, func() {
		Configure(bind.AudioSpec{})
	})
}

func TestMixer_Initialize(t *testing.T) {
	// TODO: Test Mixer Initialize
}

func TestMixer_Debug(t *testing.T) {
	// TODO: Test Mixer Debug
}

func TestMixer_Debugf(t *testing.T) {
	// TODO: Test Mixer mixDebugf
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

func TestMixer_NextOut(t *testing.T) {
	// TODO: Test Mixer NextOut
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

func TestMixer_mixCleanup(t *testing.T) {
	// TODO: Test
}

func TestMixer_getSource(t *testing.T) {
	// TODO: Test Mixer getSource
}

func TestMixer_mixVolume(t *testing.T) {
	mixChannels = 1
	assert.Equal(t, float64(0), mixVolume(0, 0, 0))
	assert.Equal(t, float64(1), mixVolume(0, 1, .5))
	mixChannels = 2
	assert.Equal(t, float64(1), mixVolume(0, 1, -.5))
	assert.Equal(t, float64(.75), mixVolume(1, 1, .5))
	assert.Equal(t, float64(.5), mixVolume(0, .5, 0))
	assert.Equal(t, float64(.5), mixVolume(1, .5, 1))
	mixChannels = 3
	assert.Equal(t, float64(1), mixVolume(0, 1, 0))
	assert.Equal(t, float64(0.6666666666666667), mixVolume(1, 1, -1))
	assert.Equal(t, float64(0.6666666666666667), mixVolume(2, .5, -.5))
	assert.Equal(t, float64(0.6666666666666667), mixVolume(1, .5, 1))
	mixChannels = 4
	assert.Equal(t, float64(1), mixVolume(0, 1, -1))
	assert.Equal(t, float64(1), mixVolume(1, 1, 0))
	assert.Equal(t, float64(.75), mixVolume(2, .5, .5))
	assert.Equal(t, float64(.625), mixVolume(3, .5, -.5))
}

// TODO: test atomix.GetSpec()

// TODO: test atomix.Debug(true) and atomix.Debug(false)

// TODO: test atomix.Play("filename", time, duration, volume)

// TODO: test sources are queued and loaded properly

// TODO: test audio sources are mixed properly into buffer

// TODO: test different timing of ^

// TODO: test different audio format / bitrate / samples of ^

// TODO: test buffer properly reported to AudioCallback
