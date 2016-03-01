// Package fire model an audio source playing at a specific time
package fire

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gopkg.in/ontomix.v0/bind/spec"
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
