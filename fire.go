/** Copyright 2015 Outright Mental, Inc. */
package atomix // is for sequence mixing

import (
	"time"
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
	begin time.Duration
	duration time.Duration
	volume float64
}

func (f *Fire) Source() string {
	return f.source
}

func (f *Fire) At(freq float64, at time.Duration) Hz {
	rel := at - f.begin
	if rel >= 0 && rel < f.duration {
      return Hz(freq * float64(at.Nanoseconds()) / float64(1000000000))
	} else {
		return 0
	}
}

/*
 *
 private */

