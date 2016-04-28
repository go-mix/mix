// Package spec specifies valid audio formats
package spec

import (
	"time"
)

// Tz is the unit of measurement of samples-over-time, e.g. for 48000Hz playback there are 48,000 Tz in 1 second.
type Tz uint64

// AudioSpec represents the frequency, format, # channels and sample rate of any audio I/O
type AudioSpec struct {
	Freq     float64
	Format   AudioFormat
	Channels int
	Length   time.Duration
}

// Validate these specs
func (spec *AudioSpec) Validate() {
	if spec.Freq == 0 {
		panic("Must specify Frequency")
	}
	if spec.Freq < 0 {
		panic("Must specify a mixing frequency greater than zero.")
	}
	if spec.Format == "" {
		panic("Must specify Format")
	}
	if spec.Channels == 0 {
		panic("Must specify Channels")
	}
}

// AudioFormat represents the bit allocation for a single sample of audio
type AudioFormat string

// AudioU8 is unsigned-integer 8-bit sample (per channel)
const AudioU8 AudioFormat = "U8"

// AudioS8 is signed-integer 8-bit sample (per channel)
const AudioS8 AudioFormat = "S8"

// AudioU16 is unsigned-integer 16-bit sample (per channel)
const AudioU16 AudioFormat = "U16"

// AudioS16 is signed-integer 16-bit sample (per channel)
const AudioS16 AudioFormat = "S16"

// AudioS32 is signed-integer 32-bit sample (per channel)
const AudioS32 AudioFormat = "S32"

// AudioF32 is floating-point 32-bit sample (per channel)
const AudioF32 AudioFormat = "F32"

// AudioF64 is floating-point 64-bit sample (per channel)
const AudioF64 AudioFormat = "F64"
