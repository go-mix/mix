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
	fire := New(src, bgnTz, endTz, vol, pan, 0, 0, 1.0, 0)
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
	// Since release is 0, it should go directly to done
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

func TestADSR_NoEnvelope(t *testing.T) {
	// Test with no ADSR (all zeros except sustain = 1.0)
	src := "sound.wav"
	bgnTz := spec.Tz(1000)
	endTz := bgnTz + spec.Tz(100)
	fire := New(src, bgnTz, endTz, 1.0, 0, 0, 0, 1.0, 0)
	
	// Envelope should always be 1.0 when no ADSR is configured
	assert.Equal(t, 1.0, fire.Envelope(bgnTz))
	assert.Equal(t, 1.0, fire.Envelope(bgnTz+50))
	assert.Equal(t, 1.0, fire.Envelope(bgnTz+99))
}

func TestADSR_AttackPhase(t *testing.T) {
	// Test attack phase
	src := "sound.wav"
	bgnTz := spec.Tz(1000)
	endTz := bgnTz + spec.Tz(200)
	attack := spec.Tz(100)
	fire := New(src, bgnTz, endTz, 1.0, 0, attack, 0, 1.0, 0)
	
	// At start, envelope should be 0
	assert.Equal(t, 0.0, fire.Envelope(bgnTz))
	// Midway through attack, should be 0.5
	assert.Equal(t, 0.5, fire.Envelope(bgnTz+50))
	// At end of attack, should be 1.0
	assert.Equal(t, 1.0, fire.Envelope(bgnTz+100))
}

func TestADSR_DecayPhase(t *testing.T) {
	// Test decay phase with sustain at 0.7
	src := "sound.wav"
	bgnTz := spec.Tz(1000)
	endTz := bgnTz + spec.Tz(300)
	attack := spec.Tz(50)
	decay := spec.Tz(100)
	sustain := 0.7
	fire := New(src, bgnTz, endTz, 1.0, 0, attack, decay, sustain, 0)
	
	// At end of attack (start of decay), should be 1.0
	assert.Equal(t, 1.0, fire.Envelope(bgnTz+attack))
	// Midway through decay, should be between 1.0 and 0.7
	midDecay := fire.Envelope(bgnTz + attack + 50)
	assert.True(t, midDecay > 0.7 && midDecay < 1.0, "Mid-decay envelope should be between sustain and peak")
	assert.InDelta(t, 0.85, midDecay, 0.01) // Should be approximately 0.85
	// At end of decay, should be sustain level (0.7)
	assert.Equal(t, sustain, fire.Envelope(bgnTz+attack+decay))
}

func TestADSR_SustainPhase(t *testing.T) {
	// Test sustain phase
	src := "sound.wav"
	bgnTz := spec.Tz(1000)
	endTz := bgnTz + spec.Tz(500)
	attack := spec.Tz(50)
	decay := spec.Tz(50)
	sustain := 0.6
	fire := New(src, bgnTz, endTz, 1.0, 0, attack, decay, sustain, 0)
	
	// During sustain phase, envelope should stay at sustain level
	assert.Equal(t, sustain, fire.Envelope(bgnTz+attack+decay))
	assert.Equal(t, sustain, fire.Envelope(bgnTz+attack+decay+100))
	assert.Equal(t, sustain, fire.Envelope(bgnTz+attack+decay+200))
}

func TestADSR_ReleasePhase(t *testing.T) {
	// Test release phase
	src := "sound.wav"
	bgnTz := spec.Tz(1000)
	endTz := bgnTz + spec.Tz(200)
	attack := spec.Tz(20)
	decay := spec.Tz(30)
	sustain := 0.8
	release := spec.Tz(50)
	fire := New(src, bgnTz, endTz, 1.0, 0, attack, decay, sustain, release)
	
	// Play through to end to trigger release
	fire.At(bgnTz)     // Start
	for i := spec.Tz(1); i <= 200; i++ {
		fire.At(bgnTz + i)
	}
	
	// Should now be in release state
	assert.Equal(t, fireStateRelease, fire.state)
	assert.True(t, fire.IsAlive())
	
	// At start of release, envelope should be at sustain level
	assert.Equal(t, sustain, fire.Envelope(endTz))
	
	// Midway through release, envelope should be at half sustain
	midRelease := fire.Envelope(endTz + 25)
	assert.InDelta(t, sustain*0.5, midRelease, 0.01)
	
	// At end of release, envelope should be 0
	assert.Equal(t, 0.0, fire.Envelope(endTz+release))
}

func TestADSR_FullCycle(t *testing.T) {
	// Test a complete ADSR cycle
	src := "sound.wav"
	bgnTz := spec.Tz(1000)
	endTz := bgnTz + spec.Tz(300)
	attack := spec.Tz(50)
	decay := spec.Tz(50)
	sustain := 0.7
	release := spec.Tz(50)
	fire := New(src, bgnTz, endTz, 1.0, 0, attack, decay, sustain, release)
	
	// Attack phase
	assert.Equal(t, 0.0, fire.Envelope(bgnTz))
	assert.True(t, fire.Envelope(bgnTz+25) > 0.0 && fire.Envelope(bgnTz+25) < 1.0)
	assert.Equal(t, 1.0, fire.Envelope(bgnTz+attack))
	
	// Decay phase
	decayMid := fire.Envelope(bgnTz + attack + 25)
	assert.True(t, decayMid > sustain && decayMid < 1.0)
	assert.Equal(t, sustain, fire.Envelope(bgnTz+attack+decay))
	
	// Sustain phase
	assert.Equal(t, sustain, fire.Envelope(bgnTz+attack+decay+100))
	
	// Trigger release by playing through
	fire.At(bgnTz)
	for i := spec.Tz(1); i <= 300; i++ {
		fire.At(bgnTz + i)
	}
	assert.Equal(t, fireStateRelease, fire.state)
	
	// Release phase
	assert.Equal(t, sustain, fire.Envelope(endTz))
	releaseMid := fire.Envelope(endTz + 25)
	assert.True(t, releaseMid > 0.0 && releaseMid < sustain)
	assert.Equal(t, 0.0, fire.Envelope(endTz+release))
}

func TestADSR_ReleaseTransition(t *testing.T) {
	// Test transition from play to release state
	src := "sound.wav"
	bgnTz := spec.Tz(100)
	endTz := bgnTz + spec.Tz(50)
	release := spec.Tz(20)
	fire := New(src, bgnTz, endTz, 1.0, 0, 0, 0, 1.0, release)
	
	// Start playing
	fire.At(bgnTz)
	assert.Equal(t, fireStatePlay, fire.state)
	
	// Play through most of the sound
	for i := spec.Tz(1); i < 50; i++ {
		fire.At(bgnTz + i)
	}
	assert.Equal(t, fireStatePlay, fire.state)
	
	// At endTz, should transition to release
	fire.At(endTz)
	assert.Equal(t, fireStateRelease, fire.state)
	assert.True(t, fire.IsAlive())
	
	// During release, should still be alive
	for i := spec.Tz(1); i < release; i++ {
		fire.At(endTz + i)
		assert.True(t, fire.IsAlive())
	}
	
	// After release completes, should be done
	fire.At(endTz + release)
	assert.Equal(t, fireStateDone, fire.state)
	assert.False(t, fire.IsAlive())
}

func TestADSR_EdgeCases(t *testing.T) {
	// Test edge cases for envelope calculation
	src := "sound.wav"
	bgnTz := spec.Tz(1000)
	endTz := bgnTz + spec.Tz(200)
	attack := spec.Tz(50)
	decay := spec.Tz(50)
	sustain := 0.7
	release := spec.Tz(50)
	fire := New(src, bgnTz, endTz, 1.0, 0, attack, decay, sustain, release)
	
	// Before begin time, envelope should be 0
	assert.Equal(t, 0.0, fire.Envelope(bgnTz-10))
	assert.Equal(t, 0.0, fire.Envelope(bgnTz-1))
	
	// Trigger release
	fire.At(bgnTz)
	for i := spec.Tz(1); i <= 200; i++ {
		fire.At(bgnTz + i)
	}
	assert.Equal(t, fireStateRelease, fire.state)
	
	// Past end of release, envelope should be 0
	assert.Equal(t, 0.0, fire.Envelope(endTz+release))
	assert.Equal(t, 0.0, fire.Envelope(endTz+release+10))
	assert.Equal(t, 0.0, fire.Envelope(endTz+release+100))
}
