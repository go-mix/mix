// Package mix combines sources into an output audio stream
package mix

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gopkg.in/mix.v0/bind/spec"
	"time"
)

//
// Tests
//

func TestBase(t *testing.T) {
	Configure(spec.AudioSpec{
		Freq:     44100,
		Format:   spec.AudioU16,
		Channels: 2,
	})
	assert.NotNil(t, Spec())
}

func TestRequiresProperAudioSpec(t *testing.T) {
	assert.Panics(t, func() {
		Configure(spec.AudioSpec{})
	})
}

func TestInitialize(t *testing.T) {
	// TODO: Test Mixer Initialize
}

func TestDebug(t *testing.T) {
	// TODO: Test Mixer Debug
}

func TestDebugf(t *testing.T) {
	// TODO: Test Mixer debug.Printf
}

func TestStart(t *testing.T) {
	// TODO: Test Mixer Start
}

func TestStartAt(t *testing.T) {
	// TODO: Test Mixer StartAt
}

func TestGetStartTime(t *testing.T) {
	// TODO: Test Mixer GetStartTime
}

func TestSetFire(t *testing.T) {
	// TODO: Test Mixer SetFire
}

func TestSetSoundsPath(t *testing.T) {
	// TODO: Test Mixer SetSoundsPath
}

func TestNextOut(t *testing.T) {
	// TODO: Test Mixer NextOut
}

func TestTeardown(t *testing.T) {
	// TODO: Test Mixer Teardown
}

func TestNextSample(t *testing.T) {
	// TODO: Test Mixer nextSample
}

func TestOutputStart(t *testing.T) {
	// TODO: Test
}

func TestOutputContinueTo(t *testing.T) {
	// TODO: Test
}

func TestOutputClose(t *testing.T) {
	// TODO: Test
}

func TestSourceAtTz(t *testing.T) {
	// TODO: Test Mixer sourceAt
}

func TestSetSpec(t *testing.T) {
	// TODO: Test Mixer setSpec
}

func TestGetSpec(t *testing.T) {
	// TODO: Test Mixer getSpec
}

func TestPrepareSource(t *testing.T) {
	// TODO: Test Mixer prepareSource
}

func TestMixCleanup(t *testing.T) {
	// TODO: Test
}

func TestMixSetSpec(t *testing.T) {
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

func TestGetSource(t *testing.T) {
	// TODO: Test Mixer getSource
}

func TestMixCycle(t *testing.T) {
	// TODO: Test garbage collection of unused sources
	// TODO: Test garbage collection of unused fires
}

// TODO: test mix.GetSpec()

// TODO: test mix.Debug(true) and mix.Debug(false)

// TODO: test mix.Play("filename", time, duration, volume)

// TODO: test sources are queued and loaded properly

// TODO: test audio sources are mixed properly into buffer

// TODO: test different timing of ^

// TODO: test different audio format / bit rate / samples of ^

// TODO: test buffer properly reported to AudioCallback
