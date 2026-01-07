package sample

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample_New(t *testing.T) {
	// Test creating a new sample with mono audio (1 channel)
	values := []Value{0.5}
	sample := New(values)
	assert.NotNil(t, sample)
	assert.Equal(t, 1, len(sample.Values))
	assert.Equal(t, Value(0.5), sample.Values[0])
}

func TestSample_NewStereo(t *testing.T) {
	// Test creating a new sample with stereo audio (2 channels)
	values := []Value{0.5, -0.3}
	sample := New(values)
	assert.NotNil(t, sample)
	assert.Equal(t, 2, len(sample.Values))
	assert.Equal(t, Value(0.5), sample.Values[0])
	assert.Equal(t, Value(-0.3), sample.Values[1])
}

func TestSample_NewMultiChannel(t *testing.T) {
	// Test creating a new sample with multiple channels
	values := []Value{0.1, 0.2, 0.3, 0.4}
	sample := New(values)
	assert.NotNil(t, sample)
	assert.Equal(t, 4, len(sample.Values))
	for i, v := range values {
		assert.Equal(t, v, sample.Values[i])
	}
}

func TestSample_EmptyValues(t *testing.T) {
	// Test creating a sample with no values
	values := []Value{}
	sample := New(values)
	assert.NotNil(t, sample)
	assert.Equal(t, 0, len(sample.Values))
}
