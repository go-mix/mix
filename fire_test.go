// Fire models a single playback event
package atomix // is for sequence mixing
// Copyright 2015 Outright Mental, Inc.

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Fire_Base(t *testing.T) {
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
	assert.Equal(t, FIRE_READY, fire.State())
	assert.Equal(t, true, fire.IsAlive())
	// start:
	assert.Equal(t, Tz(0), fire.At(bgnTz))
	assert.Equal(t, FIRE_PLAY, fire.State())
	assert.Equal(t, true, fire.IsAlive())
	// after start / before end:
	for n := Tz(1); n < testLengthTz; n++ {
		assert.Equal(t, Tz(n), fire.At(bgnTz+n))
	}
	// end:
	assert.Equal(t, testLengthTz, fire.At(endTz))
	assert.Equal(t, FIRE_DONE, fire.State())
	assert.Equal(t, false, fire.IsAlive())
	// after end:
	assert.Equal(t, Tz(0), fire.At(endTz+1))
}
