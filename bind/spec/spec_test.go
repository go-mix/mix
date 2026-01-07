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

func TestAudioSpec_AllFields(t *testing.T) {
	// Test creating spec with all fields
	spec := AudioSpec{
		Freq:     48000,
		Format:   AudioF64,
		Channels: 5,
		Length:   time.Hour,
	}
	
	assert.Equal(t, float64(48000), spec.Freq)
	assert.Equal(t, AudioF64, spec.Format)
	assert.Equal(t, 5, spec.Channels)
	assert.Equal(t, time.Hour, spec.Length)
	
	// Should validate successfully
	assert.NotPanics(t, func() {
		spec.Validate()
	})
}

func TestAudioFormat_AllConstants(t *testing.T) {
	// Verify all format constants have unique values
	formats := []AudioFormat{
		AudioU8, AudioS8, AudioU16, AudioS16,
		AudioS32, AudioF32, AudioF64,
	}
	
	// Check they're all different
	seen := make(map[AudioFormat]bool)
	for _, format := range formats {
		assert.False(t, seen[format], "Format should be unique: %v", format)
		seen[format] = true
	}
	
	// Verify count
	assert.Equal(t, 7, len(formats))
}
