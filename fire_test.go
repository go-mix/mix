/** Copyright 2015 Outright Mental, Inc. */
package atomix // is for sequence mixing

// TODO: test fire instantiates with time, velocity, duration

// TODO: test fire fires

import (
	"github.com/stretchr/testify/assert"
	// "github.com/veandco/go-sdl2/sdl"
	"testing"
	"time"
)

//
// Tests
//

func TestFireAtCorrectHz(t *testing.T) {
	freq := float64(44100)
	src := "sound.wav"
	bgn := 1 * time.Second
	dur := 500 * time.Millisecond
	hzDur := time.Second / time.Duration(freq)
	vol := float64(1)
	fire := NewFire(src, bgn, dur, vol)
	t.Logf("hzDur:%+v",hzDur)
	// note: currently, the actual duration provided is ignored; still, these values ought to come out when queried in this (the expected) order
	assert.Equal(t, Hz(0), fire.NextHzAt(0))
	assert.Equal(t, Hz(0), fire.NextHzAt(bgn - hzDur))
	assert.Equal(t, Hz(0), fire.NextHzAt(bgn))
	assert.Equal(t, Hz(1), fire.NextHzAt(bgn + hzDur))
	assert.Equal(t, Hz(2), fire.NextHzAt(bgn + 2 * hzDur))
}

	// fmt.Printf("f.source:%v at:%v f.begin:%v rel:%v\n", f.source, at, f.begin, rel)

	// Configure(sdl.AudioSpec{
	// 	Freq:     44100,
	// 	Format:   sdl.AUDIO_U16,
	// 	Channels: 2,
	// 	Samples:  4096,
	// })
	// assert.NotNil(t, Spec())
