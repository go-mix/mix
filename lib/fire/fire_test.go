// Package fire model an audio source playing at a specific time
package fire

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-mix/mix/bind/spec"
)

func TestBase(t *testing.T) {
	testLengthTz := spec.Tz(100)
	src := "sound.wav"
	bgnTz := spec.Tz(5984)
	endTz := bgnTz + testLengthTz
	vol := float64(1)
	pan := float64(0)
	pitch := float64(1.0)      // no pitch shift
	timeStretch := float64(1.0) // no time stretch
	fire := New(src, bgnTz, endTz, vol, pan, pitch, timeStretch)
	// before start:
	assert.Equal(t, spec.Tz(0), fire.At(bgnTz-2))
	assert.Equal(t, spec.Tz(0), fire.At(bgnTz-1))
	assert.Equal(t, fireStateReady, fire.state)
	assert.Equal(t, true, fire.IsAlive())
	// start:
	assert.Equal(t, spec.Tz(0), fire.At(bgnTz))
	assert.Equal(t, fireStatePlay, fire.state)
	assert.Equal(t, true, fire.IsAlive())
	// after start / before end:
	for n := spec.Tz(1); n < testLengthTz; n++ {
		assert.Equal(t, spec.Tz(n), fire.At(bgnTz+n))
	}
	// end:
	assert.Equal(t, testLengthTz, fire.At(endTz))
	assert.Equal(t, fireStateDone, fire.state)
	assert.Equal(t, false, fire.IsAlive())
	// after end:
	assert.Equal(t, spec.Tz(0), fire.At(endTz+1))
}

func TestNewFire(t *testing.T) {
	// TODO
}

func TestAt(t *testing.T) {
	// TODO
}

func TestState(t *testing.T) {
	// TODO
}

func TestIsAlive(t *testing.T) {
	// TODO
}

func TestIsPlaying(t *testing.T) {
	// TODO
}

func TestSetState(t *testing.T) {
	// TODO
}

func TestSourceLength(t *testing.T) {
	// TODO
}

func TestTeardown(t *testing.T) {
	// TODO
}

func TestPitchShift(t *testing.T) {
	src := "sound.wav"
	bgnTz := spec.Tz(5984)
	vol := float64(1)
	pan := float64(0)
	pitch := float64(2.0) // up one octave (twice as fast playback)
	timeStretch := float64(1.0)
	fire := New(src, bgnTz, 0, vol, pan, pitch, timeStretch)
	
	// Test that pitch shift affects playback position
	fire.At(bgnTz) // start - transitions to play state, initializes PlaybackTz to pitch value
	assert.Equal(t, fireStatePlay, fire.state)
	assert.Equal(t, float64(2.0), fire.PlaybackTz) // PlaybackTz initialized to pitch
	
	// With pitch=2.0, each At() call should advance by 2 samples in source
	firstSample := fire.At(bgnTz + 1)
	assert.Equal(t, spec.Tz(2), firstSample) // Should return sample 2
	assert.Equal(t, float64(4.0), fire.PlaybackTz) // PlaybackTz advances by pitch to 4.0
	
	secondSample := fire.At(bgnTz + 2)
	assert.Equal(t, spec.Tz(4), secondSample) // Should return sample 4
	assert.Equal(t, float64(6.0), fire.PlaybackTz) // PlaybackTz advances to 6.0
}

func TestTimeStretch(t *testing.T) {
	// TODO: Implement time stretching test when feature is complete
}

