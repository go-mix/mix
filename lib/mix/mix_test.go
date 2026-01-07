// Package mix combines sources into an output audio stream
package mix

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-mix/mix/bind/spec"
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

func TestMixingAlgorithmDefault(t *testing.T) {
	Configure(spec.AudioSpec{
		Freq:     44100,
		Format:   spec.AudioU16,
		Channels: 2,
	})
	// Should default to logarithmic
	assert.Equal(t, spec.MixLogarithmic, mixAlgorithm)
}

func TestMixingAlgorithmLogarithmic(t *testing.T) {
	Configure(spec.AudioSpec{
		Freq:      44100,
		Format:    spec.AudioU16,
		Channels:  2,
		Algorithm: spec.MixLogarithmic,
	})
	assert.Equal(t, spec.MixLogarithmic, mixAlgorithm)
}

func TestMixingAlgorithmLinear(t *testing.T) {
	Configure(spec.AudioSpec{
		Freq:      44100,
		Format:    spec.AudioU16,
		Channels:  2,
		Algorithm: spec.MixLinear,
	})
	assert.Equal(t, spec.MixLinear, mixAlgorithm)
}

func TestMixLinearClamp(t *testing.T) {
	// Test clamping at boundaries
	assert.Equal(t, float32(-1.0), float32(mixLinearClamp(-2.5)))
	assert.Equal(t, float32(-1.0), float32(mixLinearClamp(-1.5)))
	assert.Equal(t, float32(1.0), float32(mixLinearClamp(2.5)))
	assert.Equal(t, float32(1.0), float32(mixLinearClamp(1.5)))

	// Test pass-through in range
	assert.Equal(t, float32(0.0), float32(mixLinearClamp(0.0)))
	assert.Equal(t, float32(0.5), float32(mixLinearClamp(0.5)))
	assert.Equal(t, float32(-0.5), float32(mixLinearClamp(-0.5)))
	assert.Equal(t, float32(1.0), float32(mixLinearClamp(1.0)))
	assert.Equal(t, float32(-1.0), float32(mixLinearClamp(-1.0)))
}

func TestMixLogarithmicRangeCompression(t *testing.T) {
	// Test that values in range are scaled by golden ratio
	result := mixLogarithmicRangeCompression(0.5)
	expected := 0.5 / 1.61803398875
	assert.InDelta(t, expected, float64(result), 0.0001)

	// Test that values outside range are compressed logarithmically
	// Values > 1 should be compressed
	resultAbove := mixLogarithmicRangeCompression(2.0)
	assert.True(t, resultAbove > 0 && resultAbove < 1.5)

	// Values < -1 should be compressed
	resultBelow := mixLogarithmicRangeCompression(-2.0)
	assert.True(t, resultBelow < 0 && resultBelow > -1.5)
}

func TestMixApplyAlgorithm(t *testing.T) {
	// Test with linear algorithm
	Configure(spec.AudioSpec{
		Freq:      44100,
		Format:    spec.AudioU16,
		Channels:  2,
		Algorithm: spec.MixLinear,
	})
	result := mixApplyAlgorithm(2.0)
	assert.Equal(t, float32(1.0), float32(result))

	// Test with logarithmic algorithm
	Configure(spec.AudioSpec{
		Freq:      44100,
		Format:    spec.AudioU16,
		Channels:  2,
		Algorithm: spec.MixLogarithmic,
	})
	result = mixApplyAlgorithm(0.5)
	expected := 0.5 / 1.61803398875
	assert.InDelta(t, expected, float64(result), 0.0001)
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
