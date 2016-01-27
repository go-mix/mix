// Fire models the playback of a source
package atomix // is for sequence mixing
// Copyright 2015 Outright Mental, Inc.

import ()

func NewFire(source string, beginTz Tz, endTz Tz, volume float64, pan float64) *Fire {
	// Debugf("NewFire(%v, %v, %v, %v, %v)\n", source, beginTz, endTz, volume, pan)
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
	//	Debugf("*Fire[%s].At(%v vs %v)\n", f.Source, at, f.BeginTz)
	switch f.state {
	case FIRE_READY:
		if at >= f.BeginTz {
			f.state = FIRE_PLAY
			f.nowTz++
		}
	case FIRE_PLAY:
		t = f.nowTz
		f.nowTz++
		if f.EndTz != 0 {
			if at >= f.EndTz {
				f.state = FIRE_DONE
			}
		} else {
			f.EndTz = f.BeginTz + f.SourceLength()
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

func (f *Fire) IsPlaying() bool {
	return f.state == FIRE_PLAY
}

func (f *Fire) SetState(state FireStateEnum) {
	f.state = state
}

func (f *Fire) SourceLength() Tz {
	// TODO: evaluate if this is a bad circular dependency to call the singleton from here?
	return mixSourceLength(f.Source)
}

func (f *Fire) Teardown() {
	// TODO: confirm that all memory of this object is released when its pointer is deleted from the *Mixer.fires slice, else make sure it does get released somehow
}

/*
 *
 private */

type FireStateEnum uint

const (
	FIRE_READY FireStateEnum = 1
	FIRE_PLAY  FireStateEnum = 2
	// it is assumed that all alive states are < SOURCE_FINISHED
	FIRE_DONE FireStateEnum = 6
)
