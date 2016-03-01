// Package sample models an audio sample
package sample

import (
	"encoding/binary"
	"math"
)

func FromByteU8(sample byte) float64 {
	return float64(int8(sample))/float64(0x7F) - float64(1)
}

func FromByteS8(sample byte) float64 {
	return float64(int8(sample)) / float64(0x7F)
}

func FromBytesU16LSB(sample []byte) float64 {
	return float64(binary.LittleEndian.Uint16(sample))/float64(0x8000) - float64(1)
}

//func FromBytesU16MSB(sample []byte) float64 {
//	return float64(binary.BigEndian.Uint16(sample))/float64(0x8000) - float64(1)
//}

func FromBytesS16LSB(sample []byte) float64 {
	return float64(int16(binary.LittleEndian.Uint16(sample))) / float64(0x7FFF)
}

//func FromBytesS16MSB(sample []byte) float64 {
//	return float64(int16(binary.BigEndian.Uint16(sample))) / float64(0x7FFF)
//}

func FromBytesS32LSB(sample []byte) float64 {
	return float64(int32(binary.LittleEndian.Uint32(sample))) / float64(0x7FFFFFFF)
}

//func FromBytesS32MSB(sample []byte) float64 {
//	return float64(int32(binary.BigEndian.Uint32(sample))) / float64(0x7FFFFFFF)
//}

func FromBytesF32LSB(sample []byte) float64 {
	return float64(math.Float32frombits(binary.LittleEndian.Uint32(sample)))
}

//func FromBytesF32MSB(sample []byte) float64 {
//	return float64(math.Float32frombits(binary.BigEndian.Uint32(sample)))
//}

func FromBytesF64LSB(sample []byte) float64 {
	return float64(math.Float64frombits(binary.LittleEndian.Uint64(sample)))
}

//func FromBytesF64MSB(sample []byte) float64 {
//	return float64(math.Float64frombits(binary.BigEndian.Uint64(sample)))
//}
