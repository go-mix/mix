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
	fire := New(src, bgnTz, endTz, vol, pan)
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
	src := "test.wav"
	bgnTz := spec.Tz(1000)
	endTz := spec.Tz(2000)
	vol := float64(0.8)
	pan := float64(-0.5)
	
	fire := New(src, bgnTz, endTz, vol, pan)
	
	assert.NotNil(t, fire)
	assert.Equal(t, src, fire.Source)
	assert.Equal(t, bgnTz, fire.BeginTz)
	assert.Equal(t, endTz, fire.EndTz)
	assert.Equal(t, vol, fire.Volume)
	assert.Equal(t, pan, fire.Pan)
	assert.Equal(t, fireStateReady, fire.state)
}

func TestAt(t *testing.T) {
	src := "test.wav"
	bgnTz := spec.Tz(100)
	endTz := spec.Tz(200)
	fire := New(src, bgnTz, endTz, 1.0, 0)
	
	// Before begin: should return 0 and stay in ready state
	result := fire.At(bgnTz - 1)
	assert.Equal(t, spec.Tz(0), result)
	assert.Equal(t, fireStateReady, fire.state)
	
	// At begin: should transition to play state
	result = fire.At(bgnTz)
	assert.Equal(t, spec.Tz(0), result)
	assert.Equal(t, fireStatePlay, fire.state)
	
	// During playback: should increment
	result = fire.At(bgnTz + 1)
	assert.Equal(t, spec.Tz(1), result)
	
	result = fire.At(bgnTz + 2)
	assert.Equal(t, spec.Tz(2), result)
}

func TestState(t *testing.T) {
	src := "test.wav"
	bgnTz := spec.Tz(100)
	endTz := spec.Tz(150)
	fire := New(src, bgnTz, endTz, 1.0, 0)
	
	// Initial state should be ready
	assert.Equal(t, fireStateReady, fire.state)
	
	// After reaching begin time, state should be play
	fire.At(bgnTz)
	assert.Equal(t, fireStatePlay, fire.state)
	
	// After reaching end time, state should be done
	fire.At(endTz)
	assert.Equal(t, fireStateDone, fire.state)
}

func TestIsAlive(t *testing.T) {
	src := "test.wav"
	bgnTz := spec.Tz(100)
	endTz := spec.Tz(150)
	fire := New(src, bgnTz, endTz, 1.0, 0)
	
	// Should be alive in ready state
	assert.True(t, fire.IsAlive())
	
	// Should be alive in play state
	fire.At(bgnTz)
	assert.True(t, fire.IsAlive())
	
	// Should not be alive in done state
	fire.At(endTz)
	assert.False(t, fire.IsAlive())
}

func TestIsPlaying(t *testing.T) {
	src := "test.wav"
	bgnTz := spec.Tz(100)
	endTz := spec.Tz(150)
	fire := New(src, bgnTz, endTz, 1.0, 0)
	
	// Should not be playing in ready state
	assert.False(t, fire.IsPlaying())
	
	// Should be playing in play state
	fire.At(bgnTz)
	assert.True(t, fire.IsPlaying())
	
	// Should not be playing in done state
	fire.At(endTz)
	assert.False(t, fire.IsPlaying())
}

func TestSetState(t *testing.T) {
	// Note: there is no SetState method in the Fire struct,
	// state is managed internally through the At() method
	// This test verifies state transitions happen correctly
	src := "test.wav"
	bgnTz := spec.Tz(100)
	endTz := spec.Tz(150)
	fire := New(src, bgnTz, endTz, 1.0, 0)
	
	// Verify state transitions through At() method
	assert.Equal(t, fireStateReady, fire.state)
	fire.At(bgnTz)
	assert.Equal(t, fireStatePlay, fire.state)
	fire.At(endTz)
	assert.Equal(t, fireStateDone, fire.state)
}

func TestSourceLength(t *testing.T) {
	// sourceLength is a private method that calls source.GetLength
	// We test it indirectly through the Fire behavior when EndTz is 0
	src := "test.wav"
	bgnTz := spec.Tz(100)
	fire := New(src, bgnTz, 0, 1.0, 0) // EndTz = 0
	
	// When we call At() during play state with EndTz=0,
	// it should compute EndTz from source length
	fire.At(bgnTz) // Start playing
	fire.At(bgnTz + 1) // This should trigger EndTz calculation
	
	// EndTz should now be set (though we can't predict exact value without source)
	// We just verify it's been set to something non-zero
	assert.True(t, fire.EndTz >= bgnTz)
}

func TestTeardown(t *testing.T) {
	src := "test.wav"
	bgnTz := spec.Tz(100)
	endTz := spec.Tz(150)
	fire := New(src, bgnTz, endTz, 1.0, 0)
	
	// Teardown should not panic
	assert.NotPanics(t, func() {
		fire.Teardown()
	})
}

func TestFire_VolumeAndPan(t *testing.T) {
	src := "test.wav"
	bgnTz := spec.Tz(100)
	endTz := spec.Tz(150)
	
	// Test with different volume levels
	fire1 := New(src, bgnTz, endTz, 0.5, 0)
	assert.Equal(t, 0.5, fire1.Volume)
	
	fire2 := New(src, bgnTz, endTz, 1.0, 0)
	assert.Equal(t, 1.0, fire2.Volume)
	
	// Test with different pan values
	fireLeft := New(src, bgnTz, endTz, 1.0, -1.0)
	assert.Equal(t, -1.0, fireLeft.Pan)
	
	fireRight := New(src, bgnTz, endTz, 1.0, 1.0)
	assert.Equal(t, 1.0, fireRight.Pan)
	
	fireCenter := New(src, bgnTz, endTz, 1.0, 0.0)
	assert.Equal(t, 0.0, fireCenter.Pan)
}

func TestFire_ZeroEndTz(t *testing.T) {
	src := "test.wav"
	bgnTz := spec.Tz(100)
	
	// Create fire with EndTz = 0 (should be calculated from source)
	fire := New(src, bgnTz, 0, 1.0, 0)
	assert.Equal(t, spec.Tz(0), fire.EndTz)
	
	// Start playing
	fire.At(bgnTz)
	assert.Equal(t, fireStatePlay, fire.state)
	
	// After calling At() during play, EndTz should be calculated
	fire.At(bgnTz + 1)
	// EndTz should now be set (we can't predict exact value without source)
	assert.True(t, fire.EndTz >= bgnTz)
}

func TestFire_MultipleAtCalls(t *testing.T) {
	src := "test.wav"
	bgnTz := spec.Tz(100)
	endTz := spec.Tz(110)
	fire := New(src, bgnTz, endTz, 1.0, 0)
	
	// Multiple calls before begin should return 0
	for i := 0; i < 5; i++ {
		result := fire.At(bgnTz - 10)
		assert.Equal(t, spec.Tz(0), result)
		assert.Equal(t, fireStateReady, fire.state)
	}
	
	// Start playing
	fire.At(bgnTz)
	assert.Equal(t, fireStatePlay, fire.state)
	
	// Multiple calls during play should increment
	for i := spec.Tz(1); i <= 5; i++ {
		result := fire.At(bgnTz + i)
		assert.Equal(t, i, result)
	}
	
	// Reach end
	fire.At(endTz)
	assert.Equal(t, fireStateDone, fire.state)
	
	// Multiple calls after end should return 0
	for i := 0; i < 5; i++ {
		result := fire.At(endTz + 10)
		assert.Equal(t, spec.Tz(0), result)
		assert.Equal(t, fireStateDone, fire.state)
	}
}
