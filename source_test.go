/** Copyright 2015 Outright Mental, Inc. */
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
	assert.Equal(t, sampleBytesS16LSB([]byte{2, 0}), float64(0.24207281716360973))
	assert.Equal(t, sampleBytesS16LSB([]byte{3, 0}), float64(0.24149296548356577))
	assert.Equal(t, sampleBytesS16LSB([]byte{7, 0}), float64(0.24100466933194983))
	assert.Equal(t, sampleBytesS16LSB([]byte{7, 0}), float64(0.24069948423718984))
	assert.Equal(t, sampleBytesS16LSB([]byte{3, 0}), float64(0.2402722251045259))
	assert.Equal(t, sampleBytesS16LSB([]byte{2, 0}), float64(0.23993652150028993))
	assert.Equal(t, sampleBytesS16LSB([]byte{2, 0}), float64(0.23932615131076998))
	assert.Equal(t, sampleBytesS16LSB([]byte{5, 0}), float64(0.239112521744438))
	assert.Equal(t, sampleBytesS16LSB([]byte{1, 0}), float64(0.23838007751701407))
	assert.Equal(t, sampleBytesS16LSB([]byte{8, 5}), float64(-0.23767815179906612))
	assert.Equal(t, sampleBytesS16LSB([]byte{5, 5}), float64(-0.23746452223273415))
	assert.Equal(t, sampleBytesS16LSB([]byte{2, 5}), float64(-0.23725089266640217))
	assert.Equal(t, sampleBytesS16LSB([]byte{2, 5}), float64(-0.2369457075716422))
	assert.Equal(t, sampleBytesS16LSB([]byte{1, 5}), float64(-0.23667104098635822))
	assert.Equal(t, sampleBytesS16LSB([]byte{3, 5}), float64(-0.23630481887264626))
	assert.Equal(t, sampleBytesS16LSB([]byte{7, 5}), float64(-0.23618274483474228))
	assert.Equal(t, sampleBytesS16LSB([]byte{1, 5}), float64(-0.2357554857020783))
	assert.Equal(t, sampleBytesS16LSB([]byte{3, 5}), float64(-0.23538926358836634))
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

