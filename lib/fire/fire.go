// Package fire model an audio source playing at a specific time
package fire

import (
	"github.com/go-mix/mix/bind/spec"

	"github.com/go-mix/mix/lib/source"
)

// New Fire to represent a single audio source playing at a specific time in the future.
func New(source string, beginTz spec.Tz, endTz spec.Tz, volume float64, pan float64, pitch float64, timeStretch float64) *Fire {
	// debug.Printf("NewFire(%v, %v, %v, %v, %v)\n", source, beginTz, endTz, volume, pan)
	s := &Fire{
		/* setup */
		Source:      source,
		Volume:      volume,
		Pan:         pan,
		BeginTz:     beginTz,
		EndTz:       endTz,
		Pitch:       pitch,
		TimeStretch: timeStretch,
		/* playback */
		state:      fireStateReady,
		playbackTz: 0,
	}
	return s
}

// Fire represents a single audio source playing at a specific time in the future.
type Fire struct {
	/* setup */
	BeginTz     spec.Tz
	EndTz       spec.Tz
	Source      string
	Volume      float64 // 0 to 1
	Pan         float64 // -1 to +1
	Pitch       float64 // pitch shift multiplier (1.0 = no shift, 2.0 = up one octave, 0.5 = down one octave)
	TimeStretch float64 // time stretch multiplier (1.0 = no stretch, 2.0 = twice as slow, 0.5 = twice as fast)
	/* playback */
	nowTz      spec.Tz
	playbackTz float64 // fractional position for pitch/time stretch
	state      fireStateEnum
}

// At the series of Tz it's playing for, return the series of Tz corresponding to source audio.
func (f *Fire) At(at spec.Tz) (t spec.Tz) {
	//	debug.Printf("*Fire[%s].At(%v vs %v)\n", f.Source, at, f.BeginTz)
	switch f.state {
	case fireStateReady:
		if at >= f.BeginTz {
			f.state = fireStatePlay
			f.nowTz++
			// Initialize playbackTz based on pitch
			if f.Pitch != 0 && f.Pitch != 1.0 {
				f.playbackTz = f.Pitch
			} else {
				f.playbackTz = 1.0
			}
		}
	case fireStatePlay:
		t = spec.Tz(f.playbackTz)
		// Advance playback position based on pitch (affects playback rate)
		if f.Pitch != 0 && f.Pitch != 1.0 {
			f.playbackTz += f.Pitch
		} else {
			f.playbackTz += 1.0
		}
		f.nowTz++
		if f.EndTz != 0 {
			if at >= f.EndTz {
				f.state = fireStateDone
			}
		} else {
			actualLength := f.sourceLength()
			// Adjust end time based on pitch (faster pitch = shorter duration)
			if f.Pitch != 0 && f.Pitch != 1.0 {
				f.EndTz = f.BeginTz + spec.Tz(float64(actualLength)/f.Pitch)
			} else {
				f.EndTz = f.BeginTz + actualLength
			}
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

// Teardown the Fire and release its memory
func (f *Fire) Teardown() {
	// TODO: confirm that all memory of this object is released when its pointer is deleted from the *Mixer.fires slice, else make sure it does get released somehow
}

//
// Private
//

type fireStateEnum uint

const (
	fireStateReady fireStateEnum = 1
	fireStatePlay  fireStateEnum = 2
	// it is assumed that all alive states are < SOURCE_FINISHED
	fireStateDone fireStateEnum = 6
)

func (f *Fire) sourceLength() spec.Tz {
	return source.GetLength(f.Source)
}
