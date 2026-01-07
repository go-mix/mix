// Package source models a single audio source
package source

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-mix/mix/bind/debug"
	"github.com/go-mix/mix/bind/sample"
	"github.com/go-mix/mix/bind/spec"
)

// TODO: test multi-channel source audio files

func TestBase(t *testing.T) {
	// TODO: Test Source Base
}

func TestLoad_IntVsFloat(t *testing.T) {
	debug.Configure(true)
	testSourceSetup(44100, 1)
	sourceFloat := New("testdata/Float32bitLittleEndian48000HzEstéreo.wav")
	assert.NotNil(t, sourceFloat)
	assert.Equal(t, spec.AudioF32, sourceFloat.Spec().Format)
	sourceInt := New("testdata/Signed16bitLittleEndian44100HzMono.wav")
	assert.NotNil(t, sourceInt)
	assert.Equal(t, spec.AudioS16, sourceInt.Spec().Format)
}

func TestLoad_FAIL(t *testing.T) {
	pathFail := "testdata/ThisShouldFailBecauseItDoesNotExist.wav"
	defer func() {
		msg := recover()
		assert.IsType(t, "", msg)
		assert.Equal(t, "File not found: "+pathFail, msg)
	}()
	debug.Configure(true)
	testSourceSetup(44100, 1)
	source := New(pathFail)
	assert.NotNil(t, source)
}

func TestLoadSigned16bitLittleEndian44100HzMono(t *testing.T) {
	debug.Configure(true)
	testSourceSetup(44100, 1)
	source := New("testdata/Signed16bitLittleEndian44100HzMono.wav")
	assert.NotNil(t, source)
	totalSoundMovement := testSourceAssertSound(t, source, 1)
	assert.True(t, totalSoundMovement > .001)
}

func TestLoadFloat32bitLittleEndian48000HzEstéreo(t *testing.T) {
	debug.Configure(true)
	testSourceSetup(48000, 2)
	source := New("testdata/Float32bitLittleEndian48000HzEstéreo.wav")
	assert.NotNil(t, source)
	totalSoundMovement := testSourceAssertSound(t, source, 2)
	assert.True(t, totalSoundMovement > .001)
}

func TestOutput(t *testing.T) {
	// TODO: Test Source plays audio
}

func TestSampleAt(t *testing.T) {
	// TODO: Test Source SampleAt
}

func TestSampleAtInterpolated(t *testing.T) {
	Configure(spec.AudioSpec{
		Freq:     44100,
		Format:   spec.AudioF32,
		Channels: 1,
	})
	source := New("testdata/Signed16bitLittleEndian44100HzMono.wav")
	assert.NotNil(t, source)
	
	// Test interpolated sampling at fractional positions
	sample1 := source.SampleAtInterpolated(1.0, 1.0, 0)
	assert.NotNil(t, sample1)
	
	// Test interpolation between samples
	sample2 := source.SampleAtInterpolated(1.5, 1.0, 0)
	assert.NotNil(t, sample2)
	
	// Interpolated value should be between the two adjacent samples
	sampleBefore := source.SampleAt(1, 1.0, 0)
	sampleAfter := source.SampleAt(2, 1.0, 0)
	
	// The interpolated value at 1.5 should be between sample at 1 and sample at 2
	for c := 0; c < len(sample2); c++ {
		if sampleBefore[c] != sampleAfter[c] {
			// Check that interpolated value is between the two samples (or equal to one if they're the same)
			minVal := sampleBefore[c]
			maxVal := sampleAfter[c]
			if minVal > maxVal {
				minVal, maxVal = maxVal, minVal
			}
			assert.True(t, sample2[c] >= minVal && sample2[c] <= maxVal, 
				"Interpolated sample should be between adjacent samples")
		}
	}
}


func TestState(t *testing.T) {
	// TODO: Test Source State
}

func TestStateName(t *testing.T) {
	// TODO: Test Source StateName
}

func TestLength(t *testing.T) {
	// TODO: Test Source reports length
}

func TestTeardown(t *testing.T) {
	// TODO: Test Source Teardown
}

func TestMixer_mixVolume(t *testing.T) {
	masterChannelsFloat = 1
	assert.Equal(t, sample.Value(0), volume(0, 0, 0))
	assert.Equal(t, sample.Value(1), volume(0, 1, .5))
	masterChannelsFloat = 2
	assert.Equal(t, sample.Value(1), volume(0, 1, -.5))
	assert.Equal(t, sample.Value(.75), volume(1, 1, .5))
	assert.Equal(t, sample.Value(.5), volume(0, .5, 0))
	assert.Equal(t, sample.Value(.5), volume(1, .5, 1))
	masterChannelsFloat = 3
	assert.Equal(t, sample.Value(1), volume(0, 1, 0))
	assert.Equal(t, sample.Value(0.6666666666666667), volume(1, 1, -1))
	assert.Equal(t, sample.Value(0.6666666666666667), volume(2, .5, -.5))
	assert.Equal(t, sample.Value(0.6666666666666667), volume(1, .5, 1))
	masterChannelsFloat = 4
	assert.Equal(t, sample.Value(1), volume(0, 1, -1))
	assert.Equal(t, sample.Value(1), volume(1, 1, 0))
	assert.Equal(t, sample.Value(.75), volume(2, .5, .5))
	assert.Equal(t, sample.Value(.625), volume(3, .5, -.5))
}

//
// Private
//

func testSourceSetup(freq float64, channels int) {
	Configure(spec.AudioSpec{
		Freq:     freq,
		Format:   spec.AudioF32,
		Channels: channels,
	})
}

func testSourceAssertSound(t *testing.T, source *Source, channels int) (totalSoundMovement sample.Value) {
	for tz := spec.Tz(0); tz < source.Length(); tz++ {
		smp := source.SampleAt(tz, 1, 0)
		assert.Equal(t, channels, len(smp))
		for c := 0; c < channels; c++ {
			totalSoundMovement += smp[c].Abs()
		}
	}
	return
}
