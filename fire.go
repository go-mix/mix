/** Copyright 2015 Outright Mental, Inc. */
package atomix // is for sequence mixing

import (
	"time"
	// "math"
)

func NewFire(source string, begin time.Duration, duration time.Duration, volume float64) *Fire {
	s := &Fire{
		source: source,
		begin: begin,
		duration: duration,
		volume: volume,
	}
	return s
}

type Fire struct {
	source string
	nextPlaybackHz Hz
	begin time.Duration
	duration time.Duration
	volume float64
}

func (f *Fire) Source() string {
	return f.source
}

func (f *Fire) Volume() float64 {
	return f.volume
}

func (f *Fire) NextHzAt(at time.Duration) (hz Hz) {
	if at >= f.begin && at < f.begin + f.duration {
		hz = f.nextPlaybackHz
		f.nextPlaybackHz++
	}
	return
}

/*
 *
 private */

