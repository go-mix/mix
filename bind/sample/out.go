// Package sample models an audio sample
package sample

import (
	"github.com/go-ontomix/ontomix/bind/spec"
)

// OutNextCallbackFunc to stream mix out from ontomix
type OutNextCallbackFunc func() []Value

func ConfigureOutput(s spec.AudioSpec) {
	outSpec = &s
}

// SetOutNextCallback to set streaming callback function
func SetOutputCallback(fn OutNextCallbackFunc) {
	outNextCallback = fn
}

// OutNext to mix the next sample for all channels, in []float64
func OutNext() []Value {
	return outNextCallback()
}

// OutNextBytes to mix the next sample for all channels, in bytes
func OutNextBytes() (out []byte) {
	in := outNextCallback()
	for ch := 0; ch < outSpec.Channels; ch++ {
		switch outSpec.Format {
		case spec.AudioU8:
			out = append(out, in[ch].ToByteU8())
		case spec.AudioS8:
			out = append(out, in[ch].ToByteS8())
		case spec.AudioS16:
			out = append(out, in[ch].ToBytesS16LSB()...)
		case spec.AudioU16:
			out = append(out, in[ch].ToBytesU16LSB()...)
		case spec.AudioS32:
			out = append(out, in[ch].ToBytesS32LSB()...)
		case spec.AudioF32:
			out = append(out, in[ch].ToBytesF32LSB()...)
		case spec.AudioF64:
			out = append(out, in[ch].ToBytesF64LSB()...)
		}
	}
	return
}

/*
 *
 private */

var (
	outSpec         *spec.AudioSpec
	outNextCallback OutNextCallbackFunc
)
