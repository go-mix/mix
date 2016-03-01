// Package mix combines sources into an output audio stream
package mix

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-ontomix/ontomix/bind/spec"
	"time"
)

//
// Tests
//

func TestMixer_Base(t *testing.T) {
	Configure(spec.AudioSpec{
		Freq:     44100,
		Format:   spec.AudioU16,
		Channels: 2,
	})
	assert.NotNil(t, Spec())
}

func TestMixer_RequiresProperAudioSpec(t *testing.T) {
	assert.Panics(t, func() {
		Configure(spec.AudioSpec{})
	})
}

func TestMixer_Initialize(t *testing.T) {
	// TODO: Test Mixer Initialize
}

func TestMixer_Debug(t *testing.T) {
	// TODO: Test Mixer Debug
}

func TestMixer_Debugf(t *testing.T) {
	// TODO: Test Mixer debug.Printf
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

func TestMixer_mixSetSpec(t *testing.T) {
	// TODO: Test success passing in a bind.AudioSpec
	// TODO: Test sets the default mixCycleDurTz
}

func TestSetCycleDuration(t *testing.T) {
	masterFreq = 0 // simulates never having set a mix frequency
	defer func() {
		msg := recover()
		assert.IsType(t, "", msg)
		assert.Equal(t, "Must specify mixing frequency before setting cycle duration!", msg)
	}()
	SetCycleDuration(5 * time.Second)
}

func TestMixer_getSource(t *testing.T) {
	// TODO: Test Mixer getSource
}

func TestMixer_mixCycle(t *testing.T) {
	// TODO: Test garbage collection of unused sources
	// TODO: Test garbage collection of unused fires
}

// TODO: test ontomix.GetSpec()

// TODO: test ontomix.Debug(true) and ontomix.Debug(false)

// TODO: test ontomix.Play("filename", time, duration, volume)

// TODO: test sources are queued and loaded properly

// TODO: test audio sources are mixed properly into buffer

// TODO: test different timing of ^

// TODO: test different audio format / bit rate / samples of ^

// TODO: test buffer properly reported to AudioCallback
