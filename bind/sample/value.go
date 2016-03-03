// Package sample models an audio sample
package sample

import (
	"encoding/binary"
	"math"
)

type Value float64

func (this Value) Abs() Value {
	return Value(math.Abs(float64(this)))
}

func (this Value) ToByteU8() byte {
	return byte(this.ToUint8())
}

func (this Value) ToByteS8() byte {
	return byte(this.ToInt8())
}

func (this Value) ToBytesU16LSB() (out []byte) {
	out = make([]byte, 2)
	binary.LittleEndian.PutUint16(out, this.ToUint16())
	return
}

func (this Value) ToBytesS16LSB() (out []byte) {
	out = make([]byte, 2)
	binary.LittleEndian.PutUint16(out, uint16(this.ToInt16()))
	return
}

func (this Value) ToBytesS32LSB() (out []byte) {
	out = make([]byte, 4)
	binary.LittleEndian.PutUint32(out, uint32(this.ToInt32()))
	return
}

func (this Value) ToBytesF32LSB() (out []byte) {
	out = make([]byte, 4)
	binary.LittleEndian.PutUint32(out, math.Float32bits(float32(this)))
	return
}

func (this Value) ToBytesF64LSB() (out []byte) {
	out = make([]byte, 4)
	binary.LittleEndian.PutUint64(out, math.Float64bits(float64(this)))
	return
}

func (this Value) ToUint8() uint8 {
	return uint8(0x80 * (this + 1))
}

func (this Value) ToInt8() int8 {
	return int8(0x80 * this)
}

func (this Value) ToUint16() uint16 {
	return uint16(0x8000 * (this + 1))
}

func (this Value) ToInt16() int16 {
	return int16(0x8000 * this)
}

func (this Value) ToInt32() int32 {
	return int32(0x80000000 * this)
}

func ValueOfByteU8(sample byte) Value {
	return Value(int8(sample))/Value(0x7F) - Value(1)
}

func ValueOfByteS8(sample byte) Value {
	return Value(int8(sample)) / Value(0x7F)
}

func ValueOfBytesU16LSB(sample []byte) Value {
	return Value(binary.LittleEndian.Uint16(sample))/Value(0x8000) - Value(1)
}

//func ValueOfBytesU16MSB(sample []byte) Value {
//	return Value(binary.BigEndian.Uint16(sample))/Value(0x8000) - Value(1)
//}

func ValueOfBytesS16LSB(sample []byte) Value {
	return Value(int16(binary.LittleEndian.Uint16(sample))) / Value(0x7FFF)
}

//func ValueOfBytesS16MSB(sample []byte) Value {
//	return Value(int16(binary.BigEndian.Uint16(sample))) / Value(0x7FFF)
//}

func ValueOfBytesS32LSB(sample []byte) Value {
	return Value(int32(binary.LittleEndian.Uint32(sample))) / Value(0x7FFFFFFF)
}

//func ValueOfBytesS32MSB(sample []byte) Value {
//	return Value(int32(binary.BigEndian.Uint32(sample))) / Value(0x7FFFFFFF)
//}

func ValueOfBytesF32LSB(sample []byte) Value {
	return Value(math.Float32frombits(binary.LittleEndian.Uint32(sample)))
}

//func ValueOfBytesF32MSB(sample []byte) Value {
//	return Value(math.Float32frombits(binary.BigEndian.Uint32(sample)))
//}

func ValueOfBytesF64LSB(sample []byte) Value {
	return Value(math.Float64frombits(binary.LittleEndian.Uint64(sample)))
}

//func ValueOfBytesF64MSB(sample []byte) Value {
//	return Value(math.Float64frombits(binary.BigEndian.Uint64(sample)))
//}

