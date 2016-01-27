// Source imports audio for playback
package atomix // is for sequence mixing
// Copyright 2015 Outright Mental, Inc.

import (
	"github.com/stretchr/testify/assert"
	"github.com/veandco/go-sdl2/sdl"
	"testing"
)

// TODO: test multi-channel source audio files

func TestSource_Base(t *testing.T) {
	// TODO: Test Source Base
}

func TestSource_Load_FAIL(t *testing.T) {
	// TODO: Test Source loading nonexistent URL results in a graceful failure
}

func TestSource_Load(t *testing.T) {
	Debug(true)
	testSourceSetup()
	source := NewSource("./lib/S16.wav")
	assert.NotNil(t, source)
}

func TestSource_Playback(t *testing.T) {
	// TODO: Test Source plays audio
}

func TestSource_SampleAt(t *testing.T) {
	// TODO: Test Source SampleAt
}

func TestSource_State(t *testing.T) {
	// TODO: Test Source State
}

func TestSource_StateName(t *testing.T) {
	// TODO: Test Source StateName
}

func TestSource_Length(t *testing.T) {
	// TODO: Test Source reports length
}

func TestSource_Teardown(t *testing.T) {
	// TODO: Test Source Teardown
}

func TestSource_load8(t *testing.T) {
	// TODO: Test Source load8
}

func TestSource_load16(t *testing.T) {
	// TODO: Test Source load16
}

func TestSource_load32(t *testing.T) {
	// TODO: Test Source load32
}

func TestSource_sampleByteU8(t *testing.T) {
	// TODO: Test Source sampleByteU8
}

func TestSource_sampleByteS8(t *testing.T) {
	// TODO: Test Source sampleByteS8
}

func Test_sampleBytesU16LSB(t *testing.T) {
	// TODO: Test Source sampleBytesU16LSB
}

func Test_sampleBytesU16MSB(t *testing.T) {
	// TODO: Test Source sampleBytesU16MSB
}

func Test_sampleBytesS16LSB(t *testing.T) {
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

func Test_sampleBytesS16MSB(t *testing.T) {
	// TODO: Test Source sampleBytesS16MSB
}

func Test_sampleBytesS32LSB(t *testing.T) {
	// TODO: Test Source sampleBytesS32LSB
}

func Test_sampleBytesS32MSB(t *testing.T) {
	// TODO: Test Source sampleBytesS32MSB
}

func Test_sampleBytesF32LSB(t *testing.T) {
	// TODO: Test Source sampleBytesF32LSB
}

func Test_sampleBytesF32MSB(t *testing.T) {
	// TODO: Test Source sampleBytesF32MSB
}

//
// Test Components
//

func testSourceSetup() {
	Configure(sdl.AudioSpec{
		Freq:     44100,
		Format:   sdl.AUDIO_S16,
		Channels: 1,
		Samples:  4096,
	})
}
