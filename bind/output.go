// Package bind is for modular binding of atomix to audio interface
package bind

import (
	"encoding/binary"
	"math"
)

var (
	useWAV                      = OptWAVGo
	usePlayback                 = OptPlaybackSDL
	outputSpec                  *AudioSpec
	outputCallbackMixNextSample outputCallbackMixNextSampleFunc
)

type outputCallbackMixNextSampleFunc func() []float64

func outputNextBytes() (out []byte) {
	in := outputCallbackMixNextSample()
	for ch := 0; ch < outputSpec.Channels; ch++ {
		switch outputSpec.Format {
		case AudioU8:
			out = append(out, outputByteU8(in[ch]))
		case AudioS8:
			out = append(out, outputByteS8(in[ch]))
		case AudioS16:
			out = append(out, outputBytesS16LSB(in[ch])...)
		case AudioU16:
			out = append(out, outputBytesU16LSB(in[ch])...)
		case AudioS32:
			out = append(out, outputBytesS32LSB(in[ch])...)
		case AudioF32:
			out = append(out, outputBytesF32LSB(in[ch])...)
		case AudioF64:
			out = append(out, outputBytesF64LSB(in[ch])...)
		}
	}
	return
}

func outputByteU8(sample float64) byte {
	return byte(outputUint8(sample))
}

func outputByteS8(sample float64) byte {
	return byte(outputInt8(sample))
}

func outputBytesU16LSB(sample float64) (out []byte) {
	out = make([]byte, 2)
	binary.LittleEndian.PutUint16(out, outputUint16(sample))
	return
}

func outputBytesS16LSB(sample float64) (out []byte) {
	out = make([]byte, 2)
	binary.LittleEndian.PutUint16(out, uint16(outputInt16(sample)))
	return
}

func outputBytesS32LSB(sample float64) (out []byte) {
	out = make([]byte, 4)
	binary.LittleEndian.PutUint32(out, uint32(outputInt32(sample)))
	return
}

func outputBytesF32LSB(sample float64) (out []byte) {
	out = make([]byte, 4)
	binary.LittleEndian.PutUint32(out, math.Float32bits(float32(sample)))
	return
}

func outputBytesF64LSB(sample float64) (out []byte) {
	out = make([]byte, 4)
	binary.LittleEndian.PutUint64(out, math.Float64bits(sample))
	return
}

func outputUint8(sample float64) uint8 {
	return uint8(0x80 * (sample + 1))
}

func outputInt8(sample float64) int8 {
	return int8(0x80 * sample)
}

func outputUint16(sample float64) uint16 {
	return uint16(0x8000 * (sample + 1))
}

func outputInt16(sample float64) int16 {
	return int16(0x8000 * sample)
}

func outputInt32(sample float64) int32 {
	return int32(0x80000000 * sample)
}
