// Copyright 2015 Outright Mental, Inc.
package atomix // is for sequence mixing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Source_Base(t *testing.T) {
	// TODO: Test Source Base
}

func Test_Source_Load(t *testing.T) {
	Debug(true)
	source := NewSource("./lib/test.wav")
	assert.NotNil(t, source)
}

func Test_Source_Playback(t *testing.T) {
	// TODO: Test Source plays audio
}

func Test_Source_SampleAt(t *testing.T) {
	// TODO: Test Source SampleAt
}

func Test_Source_State(t *testing.T) {
	// TODO: Test Source State
}

func Test_Source_StateName(t *testing.T) {
	// TODO: Test Source StateName
}

func Test_Source_Teardown(t *testing.T) {
	// TODO: Test Source Teardown
}

func Test_Source_load8(t *testing.T) {
	// TODO: Test Source load8
}

func Test_Source_load16(t *testing.T) {
	// TODO: Test Source load16
}

func Test_Source_load32(t *testing.T) {
	// TODO: Test Source load32
}

func Test_Source_sampleByteU8(t *testing.T) {
	// TODO: Test Source sampleByteU8
}

func Test_Source_sampleByteS8(t *testing.T) {
	// TODO: Test Source sampleByteS8
}

func Test_Source_sampleBytesU16LSB(t *testing.T) {
	// TODO: Test Source sampleBytesU16LSB
}

func Test_Source_sampleBytesU16MSB(t *testing.T) {
	// TODO: Test Source sampleBytesU16MSB
}

func Test_Source_sampleBytesS16LSB(t *testing.T) {
	assert.Equal(t, sampleBytesS16LSB([]byte{51, 197}), float64(-0.4593951231421857))
	assert.Equal(t, sampleBytesS16LSB([]byte{104, 196}), float64(-0.46559038056581314))
	assert.Equal(t, sampleBytesS16LSB([]byte{160, 195}), float64(-0.47169408246101263))
	assert.Equal(t, sampleBytesS16LSB([]byte{215, 194}), float64(-0.47782830286568806))
	assert.Equal(t, sampleBytesS16LSB([]byte{24, 194}), float64(-0.4836573381756035))
	assert.Equal(t, sampleBytesS16LSB([]byte{82, 193}), float64(-0.48970000305185096))
	assert.Equal(t, sampleBytesS16LSB([]byte{142, 192}), float64(-0.4956816309091464))
	assert.Equal(t, sampleBytesS16LSB([]byte{203, 191}), float64(-0.5016327402569658))
	assert.Equal(t, sampleBytesS16LSB([]byte{14, 191}), float64(-0.5074007385479293))
	assert.Equal(t, sampleBytesS16LSB([]byte{73, 190}), float64(-0.5134128849147007))
	assert.Equal(t, sampleBytesS16LSB([]byte{132, 189}), float64(-0.5194250312814722))
	assert.Equal(t, sampleBytesS16LSB([]byte{197, 188}), float64(-0.5252540665913876))
	assert.Equal(t, sampleBytesS16LSB([]byte{5, 188}), float64(-0.5311136204107791))
	assert.Equal(t, sampleBytesS16LSB([]byte{205, 2}), float64(0.021881771294289986))
	assert.Equal(t, sampleBytesS16LSB([]byte{183, 2}), float64(0.02121036408581805))
	assert.Equal(t, sampleBytesS16LSB([]byte{171, 2}), float64(0.020844141972106083))
	assert.Equal(t, sampleBytesS16LSB([]byte{148, 2}), float64(0.020142216254158147))
	assert.Equal(t, sampleBytesS16LSB([]byte{131, 2}), float64(0.019623401593066195))
	assert.Equal(t, sampleBytesS16LSB([]byte{103, 2}), float64(0.018768883327738274))
	assert.Equal(t, sampleBytesS16LSB([]byte{94, 2}), float64(0.0184942167424543))
	assert.Equal(t, sampleBytesS16LSB([]byte{74, 2}), float64(0.017883846552934356))
	assert.Equal(t, sampleBytesS16LSB([]byte{53, 2}), float64(0.017242957853938413))
	assert.Equal(t, sampleBytesS16LSB([]byte{32, 2}), float64(0.016602069154942473))
	assert.Equal(t, sampleBytesS16LSB([]byte{14, 2}), float64(0.016052735984374525))
	assert.Equal(t, sampleBytesS16LSB([]byte{245, 1}), float64(0.015289773247474594))
	assert.Equal(t, sampleBytesS16LSB([]byte{231, 1}), float64(0.014862514114810634))
}

func Test_Source_sampleBytesS16MSB(t *testing.T) {
	// TODO: Test Source sampleBytesS16MSB
}

func Test_Source_sampleBytesS32LSB(t *testing.T) {
	// TODO: Test Source sampleBytesS32LSB
}

func Test_Source_sampleBytesS32MSB(t *testing.T) {
	// TODO: Test Source sampleBytesS32MSB
}

func Test_Source_sampleBytesF32LSB(t *testing.T) {
	// TODO: Test Source sampleBytesF32LSB
}

func Test_Source_sampleBytesF32MSB(t *testing.T) {
	// TODO: Test Source sampleBytesF32MSB
}
