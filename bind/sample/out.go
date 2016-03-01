// Package sample models an audio sample
package sample

import (
	"encoding/binary"
	"math"

	"gopkg.in/ontomix.v0/bind/spec"
)

// OutNextCallbackFunc to stream mix out from ontomix
type OutNextCallbackFunc func() []float64

func ConfigureOutput(s spec.AudioSpec) {
	outSpec = &s
}

// SetOutNextCallback to set streaming callback function
func SetOutputCallback(fn OutNextCallbackFunc) {
	outNextCallback = fn
}

// OutNext to mix the next sample for all channels, in []float64
func OutNext() []float64 {
	return outNextCallback()
}

// OutNextBytes to mix the next sample for all channels, in bytes
func OutNextBytes() (out []byte) {
	in := outNextCallback()
	for ch := 0; ch < outSpec.Channels; ch++ {
		switch outSpec.Format {
		case spec.AudioU8:
			out = append(out, outByteU8(in[ch]))
		case spec.AudioS8:
			out = append(out, outByteS8(in[ch]))
		case spec.AudioS16:
			out = append(out, outBytesS16LSB(in[ch])...)
		case spec.AudioU16:
			out = append(out, outBytesU16LSB(in[ch])...)
		case spec.AudioS32:
			out = append(out, outBytesS32LSB(in[ch])...)
		case spec.AudioF32:
			out = append(out, outBytesF32LSB(in[ch])...)
		case spec.AudioF64:
			out = append(out, outBytesF64LSB(in[ch])...)
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

func outByteU8(sample float64) byte {
	return byte(outUint8(sample))
}

func outByteS8(sample float64) byte {
	return byte(outInt8(sample))
}

func outBytesU16LSB(sample float64) (out []byte) {
	out = make([]byte, 2)
	binary.LittleEndian.PutUint16(out, outUint16(sample))
	return
}

func outBytesS16LSB(sample float64) (out []byte) {
	out = make([]byte, 2)
	binary.LittleEndian.PutUint16(out, uint16(outInt16(sample)))
	return
}

func outBytesS32LSB(sample float64) (out []byte) {
	out = make([]byte, 4)
	binary.LittleEndian.PutUint32(out, uint32(outInt32(sample)))
	return
}

func outBytesF32LSB(sample float64) (out []byte) {
	out = make([]byte, 4)
	binary.LittleEndian.PutUint32(out, math.Float32bits(float32(sample)))
	return
}

func outBytesF64LSB(sample float64) (out []byte) {
	out = make([]byte, 4)
	binary.LittleEndian.PutUint64(out, math.Float64bits(sample))
	return
}

func outUint8(sample float64) uint8 {
	return uint8(0x80 * (sample + 1))
}

func outInt8(sample float64) int8 {
	return int8(0x80 * sample)
}

func outUint16(sample float64) uint16 {
	return uint16(0x8000 * (sample + 1))
}

func outInt16(sample float64) int16 {
	return int16(0x8000 * sample)
}

func outInt32(sample float64) int32 {
	return int32(0x80000000 * sample)
}
