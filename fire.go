// Fire models a single playback event
// Copyright 2015 Outright Mental, Inc.
package atomix // is for sequence mixing

import ()

func NewFire(source string, beginTz Tz, endTz Tz, volume float64, pan float64) *Fire {
	s := &Fire{
		/* setup */
		Source:  source,
		Volume:  volume,
		Pan:     pan,
		BeginTz: beginTz,
		EndTz:   endTz,
		/* playback */
		state: FIRE_READY,
	}
	return s
}

type Fire struct {
	/* setup */
	BeginTz Tz
	EndTz   Tz
	Source  string
	Volume  float64
	Pan     float64
	/* playback */
	nowTz Tz
	state FireStateEnum
}

func (f *Fire) At(at Tz) (t Tz) {
	switch f.state {
	case FIRE_READY:
		if at >= f.BeginTz {
			f.state = FIRE_PLAY
			f.nowTz++
		}
	case FIRE_PLAY:
		t = f.nowTz
		f.nowTz++
		if at >= f.EndTz {
			f.state = FIRE_DONE
		}
	case FIRE_DONE:
		// garbage collection
	}
	return
}

func (f *Fire) State() FireStateEnum {
	return f.state
}

func (f *Fire) IsAlive() bool {
	return f.state < FIRE_DONE
}

func (f *Fire) SetState(state FireStateEnum) {
	f.state = state
}

/*
 *
 private */

type FireStateEnum uint

const (
	FIRE_READY FireStateEnum = 1
	FIRE_PLAY  FireStateEnum = 2
	// it can be assumed that all alive states are < SOURCE_FINISHED
	FIRE_DONE FireStateEnum = 6
)
