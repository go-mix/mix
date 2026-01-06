// Package sample models an audio sample
package sample

import (
"testing"

"github.com/stretchr/testify/assert"
)

func TestSampleValueRoundTrip_S16(t *testing.T) {
original := Value(0.5)
bytes := original.ToBytesS16LSB()
roundTrip := ValueOfBytesS16LSB(bytes)
diff := (original - roundTrip).Abs()
assert.True(t, diff < 0.001, "Round trip should preserve value within tolerance")
}

func TestSampleValueRoundTrip_S32(t *testing.T) {
original := Value(0.5)
bytes := original.ToBytesS32LSB()
roundTrip := ValueOfBytesS32LSB(bytes)
diff := (original - roundTrip).Abs()
assert.True(t, diff < 0.0001, "Round trip should preserve value within tolerance")
}

func TestSampleValueRoundTrip_F32(t *testing.T) {
original := Value(0.123456789)
bytes := original.ToBytesF32LSB()
roundTrip := ValueOfBytesF32LSB(bytes)
diff := (original - roundTrip).Abs()
assert.True(t, diff < 0.0001, "Round trip should preserve value within float32 tolerance")
}

func TestSampleWithMultipleValues(t *testing.T) {
values := []Value{0.1, 0.2, 0.3, 0.4, 0.5}
sample := New(values)
assert.Equal(t, len(values), len(sample.Values))
for i, v := range values {
assert.Equal(t, v, sample.Values[i])
}
}

func TestValueArithmetic(t *testing.T) {
v1 := Value(0.5)
v2 := Value(0.3)
sum := v1 + v2
assert.InDelta(t, 0.8, float64(sum), 0.0001)
diff := v1 - v2
assert.InDelta(t, 0.2, float64(diff), 0.0001)
product := v1 * v2
assert.InDelta(t, 0.15, float64(product), 0.0001)
}

func TestValueEdgeCases(t *testing.T) {
zero := Value(0)
assert.Equal(t, Value(0), zero.Abs())
tiny := Value(0.000001)
assert.True(t, tiny.Abs() > 0)
negative := Value(-0.75)
assert.Equal(t, Value(0.75), negative.Abs())
limit := Value(1.0)
assert.Equal(t, Value(1.0), limit)
assert.Equal(t, Value(1.0), limit.Abs())
}

func TestSampleCreationVariants(t *testing.T) {
mono := New([]Value{0.5})
assert.Equal(t, 1, len(mono.Values))
stereo := New([]Value{0.5, -0.3})
assert.Equal(t, 2, len(stereo.Values))
surround := New([]Value{0.1, 0.2, 0.3, 0.4, 0.5, 0.6})
assert.Equal(t, 6, len(surround.Values))
empty := New([]Value{})
assert.Equal(t, 0, len(empty.Values))
}
