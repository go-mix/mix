// Package atomix is a sequence-based Go-native audio mixer
package atomix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFire_Base(t *testing.T) {
	testLengthTz := Tz(100)
	src := "sound.wav"
	bgnTz := Tz(5984)
	endTz := bgnTz + testLengthTz
	vol := float64(1)
	pan := float64(0)
	fire := NewFire(src, bgnTz, endTz, vol, pan)
	// before start:
	assert.Equal(t, Tz(0), fire.At(bgnTz-2))
	assert.Equal(t, Tz(0), fire.At(bgnTz-1))
	assert.Equal(t, fireStateReady, fire.state)
	assert.Equal(t, true, fire.IsAlive())
	// start:
	assert.Equal(t, Tz(0), fire.At(bgnTz))
	assert.Equal(t, fireStatePlay, fire.state)
	assert.Equal(t, true, fire.IsAlive())
	// after start / before end:
	for n := Tz(1); n < testLengthTz; n++ {
		assert.Equal(t, Tz(n), fire.At(bgnTz+n))
	}
	// end:
	assert.Equal(t, testLengthTz, fire.At(endTz))
	assert.Equal(t, fireStateDone, fire.state)
	assert.Equal(t, false, fire.IsAlive())
	// after end:
	assert.Equal(t, Tz(0), fire.At(endTz+1))
}

func TestFire_NewFire(t *testing.T) {
	// TODO
}

func TestFire_At(t *testing.T) {
	// TODO
}

func TestFire_State(t *testing.T) {
	// TODO
}

func TestFire_IsAlive(t *testing.T) {
	// TODO
}

func TestFire_IsPlaying(t *testing.T) {
	// TODO
}

func TestFire_SetState(t *testing.T) {
	// TODO
}

func TestFire_SourceLength(t *testing.T) {
	// TODO
}

func TestFire_Teardown(t *testing.T) {
	// TODO
}
