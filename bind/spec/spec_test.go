// Package spec specifies valid audio formats
package spec

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAudioSpec_Validate(t *testing.T) {
	// Valid spec should not panic
	validSpec := AudioSpec{
		Freq:     44100,
		Format:   AudioS16,
		Channels: 2,
		Length:   time.Second,
	}
	assert.NotPanics(t, func() {
		validSpec.Validate()
	})
}

func TestAudioSpec_Validate_PanicOnZeroFreq(t *testing.T) {
	defer func() {
		msg := recover()
		assert.IsType(t, "", msg)
		assert.Equal(t, "Must specify Frequency", msg)
	}()
	spec := AudioSpec{
		Format:   AudioS16,
		Channels: 2,
	}
	spec.Validate()
}

func TestAudioSpec_Validate_PanicOnNegativeFreq(t *testing.T) {
	defer func() {
		msg := recover()
		assert.IsType(t, "", msg)
		assert.Equal(t, "Must specify a mixing frequency greater than zero.", msg)
	}()
	spec := AudioSpec{
		Freq:     -100,
		Format:   AudioS16,
		Channels: 2,
	}
	spec.Validate()
}

func TestAudioSpec_Validate_PanicOnEmptyFormat(t *testing.T) {
	defer func() {
		msg := recover()
		assert.IsType(t, "", msg)
		assert.Equal(t, "Must specify Format", msg)
	}()
	spec := AudioSpec{
		Freq:     44100,
		Channels: 2,
	}
	spec.Validate()
}

func TestAudioSpec_Validate_PanicOnZeroChannels(t *testing.T) {
	defer func() {
		msg := recover()
		assert.IsType(t, "", msg)
		assert.Equal(t, "Must specify Channels", msg)
	}()
	spec := AudioSpec{
		Freq:   44100,
		Format: AudioS16,
	}
	spec.Validate()
}

func TestAudioFormat_Constants(t *testing.T) {
	// Test all audio format constants are defined
	assert.Equal(t, AudioFormat("U8"), AudioU8)
	assert.Equal(t, AudioFormat("S8"), AudioS8)
	assert.Equal(t, AudioFormat("U16"), AudioU16)
	assert.Equal(t, AudioFormat("S16"), AudioS16)
	assert.Equal(t, AudioFormat("S32"), AudioS32)
	assert.Equal(t, AudioFormat("F32"), AudioF32)
	assert.Equal(t, AudioFormat("F64"), AudioF64)
}

func TestTz_Type(t *testing.T) {
	// Test that Tz can be used as expected
	var tz Tz = 48000
	assert.Equal(t, Tz(48000), tz)
	
	// Test basic arithmetic
	tz = tz * 2
	assert.Equal(t, Tz(96000), tz)
}
