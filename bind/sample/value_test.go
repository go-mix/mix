// Package sample models an audio sample
package sample

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValueFromByteU8(t *testing.T) {
	//TODO: Test
}

func TestValueFromByteS8(t *testing.T) {
	//TODO: Test
}

func TestValueFromBytesU16LSB(t *testing.T) {
	//TODO: Test
}

func TestValueFromBytesU16MSB(t *testing.T) {
	//TODO: Test
}

func TestValueFromBytesS16LSB(t *testing.T) {
	//TODO: Test
}

func TestValueFromBytesS16MSB(t *testing.T) {
	//TODO: Test
}

func TestValueFromBytesS32LSB(t *testing.T) {
	//TODO: Test
}

func TestValueFromBytesS32MSB(t *testing.T) {
	//TODO: Test
}

func TestValueFromBytesF32LSB(t *testing.T) {
	//TODO: Test
}

func TestValueFromBytesF32MSB(t *testing.T) {
	//TODO: Test
}

func TestValueFromBytesF64LSB(t *testing.T) {
	//TODO: Test
}

func TestValueFromBytesF64MSB(t *testing.T) {
	//TODO: Test
}

func TestValueToByteU8(t *testing.T) {
	// TODO
}

func TestValueToByteS8(t *testing.T) {
	// TODO
}

func TestValueToBytesU16LSB(t *testing.T) {
	// TODO
}

func TestValueToBytesS16LSB(t *testing.T) {
	// TODO
}

func TestValueToBytesS32LSB(t *testing.T) {
	// TODO
}

func TestValueToBytesF32LSB(t *testing.T) {
	// TODO
}

func TestValueToBytesF64LSB(t *testing.T) {
	// TODO
}

func TestValueToUint8(t *testing.T) {
	// TODO
}

func TestValueToInt8(t *testing.T) {
	// TODO
}

func TestValueToUint16(t *testing.T) {
	// TODO
}

func TestValueToInt16(t *testing.T) {
	// TODO
}

func TestValueToInt32(t *testing.T) {
	// TODO
}


func TestValueAbs(t *testing.T) {
// Test positive value
v1 := Value(0.5)
assert.Equal(t, Value(0.5), v1.Abs())

// Test negative value
v2 := Value(-0.5)
assert.Equal(t, Value(0.5), v2.Abs())

// Test zero
v3 := Value(0)
assert.Equal(t, Value(0), v3.Abs())

// Test edge cases
v4 := Value(-1.0)
assert.Equal(t, Value(1.0), v4.Abs())

v5 := Value(1.0)
assert.Equal(t, Value(1.0), v5.Abs())
}
