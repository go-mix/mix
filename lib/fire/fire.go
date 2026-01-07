// Package fire model an audio source playing at a specific time
package fire

import (
	"github.com/go-mix/mix/bind/spec"

	"github.com/go-mix/mix/lib/source"
)

// New Fire to represent a single audio source playing at a specific time in the future.
func New(source string, beginTz spec.Tz, endTz spec.Tz, volume float64, pan float64, attack spec.Tz, decay spec.Tz, sustain float64, release spec.Tz) *Fire {
	// debug.Printf("NewFire(%v, %v, %v, %v, %v)\n", source, beginTz, endTz, volume, pan)
	s := &Fire{
		/* setup */
		Source:  source,
		Volume:  volume,
		Pan:     pan,
		BeginTz: beginTz,
		EndTz:   endTz,
		Attack:  attack,
		Decay:   decay,
		Sustain: sustain,
		Release: release,
		/* playback */
		state: fireStateReady,
	}
	return s
}

// Fire represents a single audio source playing at a specific time in the future.
type Fire struct {
	/* setup */
	BeginTz spec.Tz
	EndTz   spec.Tz
	Source  string
	Volume  float64 // 0 to 1
	Pan     float64 // -1 to +1
	Attack  spec.Tz // ADSR: Attack time in samples
	Decay   spec.Tz // ADSR: Decay time in samples
	Sustain float64 // ADSR: Sustain level (0 to 1)
	Release spec.Tz // ADSR: Release time in samples
	/* playback */
	nowTz     spec.Tz
	releaseTz spec.Tz // Time when release phase started
	state     fireStateEnum
}

// At the series of Tz it's playing for, return the series of Tz corresponding to source audio.
func (f *Fire) At(at spec.Tz) (t spec.Tz) {
	//	debug.Printf("*Fire[%s].At(%v vs %v)\n", f.Source, at, f.BeginTz)
	switch f.state {
	case fireStateReady:
		if at >= f.BeginTz {
			f.state = fireStatePlay
			f.nowTz++
		}
	case fireStatePlay:
		t = f.nowTz
		f.nowTz++
		if f.EndTz != 0 {
			if at >= f.EndTz {
				// Start release phase
				if f.Release > 0 {
					f.releaseTz = at
					f.state = fireStateRelease
				} else {
					f.state = fireStateDone
				}
			}
		} else {
			f.EndTz = f.BeginTz + f.sourceLength()
		}
	case fireStateRelease:
		t = f.nowTz
		f.nowTz++
		// Check if release phase is complete
		if at >= f.releaseTz+f.Release {
			f.state = fireStateDone
		}
	case fireStateDone:
		// garbage collection
	}
	return
}

// IsAlive the Fire?
func (f *Fire) IsAlive() bool {
	return f.state < fireStateDone
}

// IsPlaying the Fire?
func (f *Fire) IsPlaying() bool {
	return f.state == fireStatePlay
}

// Envelope calculates the ADSR envelope multiplier at the current playback position.
// Returns a value between 0 and 1 that should be multiplied with the volume.
func (f *Fire) Envelope(at spec.Tz) float64 {
	// If the fire is done, return 0
	if f.state == fireStateDone {
		return 0.0
	}

	// If position is before start, return 0
	if at < f.BeginTz {
		return 0.0
	}

	// If no ADSR is configured and sustain is full, return 1.0 (full volume)
	if f.Attack == 0 && f.Decay == 0 && f.Sustain == 1.0 && f.Release == 0 {
		return 1.0
	}

	// Calculate position relative to start
	positionInFire := at - f.BeginTz

	// If we're in the release phase
	if f.state == fireStateRelease {
		// Check that we're past the release start time
		if at < f.releaseTz {
			// This shouldn't happen, but if it does, use sustain
			return f.Sustain
		}
		positionInRelease := at - f.releaseTz
		if f.Release > 0 {
			// Linear fade from sustain level to 0, clamped to [0, sustain]
			if positionInRelease >= f.Release {
				return 0.0
			}
			return f.Sustain * float64(f.Release-positionInRelease) / float64(f.Release)
		}
		return 0.0
	}

	// Attack phase
	if positionInFire < f.Attack {
		if f.Attack > 0 {
			return float64(positionInFire) / float64(f.Attack)
		}
		return 1.0
	}

	// Decay phase
	if positionInFire < f.Attack+f.Decay {
		if f.Decay > 0 {
			positionInDecay := positionInFire - f.Attack
			// Linear interpolation from 1.0 to sustain level
			return 1.0 - (1.0-f.Sustain)*(float64(positionInDecay)/float64(f.Decay))
		}
		return f.Sustain
	}

	// Sustain phase
	return f.Sustain
}

// Teardown the Fire and release its memory
func (f *Fire) Teardown() {
	// TODO: confirm that all memory of this object is released when its pointer is deleted from the *Mixer.fires slice, else make sure it does get released somehow
}

//
// Private
//

type fireStateEnum uint

const (
	fireStateReady   fireStateEnum = 1
	fireStatePlay    fireStateEnum = 2
	fireStateRelease fireStateEnum = 3
	// it is assumed that all alive states are < SOURCE_FINISHED
	fireStateDone fireStateEnum = 6
)

func (f *Fire) sourceLength() spec.Tz {
	return source.GetLength(f.Source)
}
