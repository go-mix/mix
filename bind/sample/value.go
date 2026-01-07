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
	out = make([]byte, 8)
	binary.LittleEndian.PutUint64(out, math.Float64bits(float64(this)))
	return
}

func (this Value) ToUint8() uint8 {
	// U8 range: 0 to 255, where 0 = -1.0, 128 = 0.0, 255 = ~1.0
	result := (float64(this) + 1.0) * 127.5 + 0.5 // Add 0.5 for rounding
	if result < 0 {
		return 0
	}
	if result > 255 {
		return 255
	}
	return uint8(result)
}

func (this Value) ToInt8() int8 {
	// S8 range: -128 to 127, where -128 = -1.0, 0 = 0.0, 127 = ~1.0
	if this <= -1.0 {
		return -128
	}
	result := math.Round(float64(this) * 127.0)
	if result > 127 {
		return 127
	}
	return int8(result)
}

func (this Value) ToUint16() uint16 {
	// U16 range: 0 to 65535, where 0 = -1.0, 32768 = 0.0, 65535 = ~1.0
	result := (float64(this) + 1.0) * 32767.5 + 0.5 // Add 0.5 for rounding
	if result < 0 {
		return 0
	}
	if result > 65535 {
		return 65535
	}
	return uint16(result)
}

func (this Value) ToInt16() int16 {
	// S16 range: -32768 to 32767, where -32768 = -1.0, 0 = 0.0, 32767 = ~1.0
	if this <= -1.0 {
		return -32768
	}
	result := math.Round(float64(this) * 32767.0)
	if result > 32767 {
		return 32767
	}
	return int16(result)
}

func (this Value) ToInt32() int32 {
	// S32 range: -2147483648 to 2147483647
	if this <= -1.0 {
		return -2147483648
	}
	result := math.Round(float64(this) * 2147483647.0)
	if result > 2147483647 {
		return 2147483647
	}
	return int32(result)
}

func ValueOfByteU8(sample byte) Value {
	// U8: 0 = -1.0, 128 = 0.0, 255 = 1.0
	// Formula: (sample / 127.5) - 1.0
	return Value(sample)/Value(127.5) - Value(1)
}

func ValueOfByteS8(sample byte) Value {
	return Value(int8(sample)) / Value(0x7F)
}

func ValueOfBytesU16LSB(sample []byte) Value {
	// U16: 0 = -1.0, 32768 = 0.0, 65535 = 1.0
	// Formula: (value / 32767.5) - 1.0
	return Value(binary.LittleEndian.Uint16(sample))/Value(32767.5) - Value(1)
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
