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
	// Test that source can provide audio samples
	testSourceSetup(44100, 1)
	testFile := "testdata/Signed16bitLittleEndian44100HzMono.wav"
	source := New(testFile)
	assert.NotNil(t, source)
	
	// Should be able to get samples
	totalMovement := testSourceAssertSound(t, source, 1)
	assert.True(t, totalMovement > 0, "Audio file should have some non-zero samples")
}

func TestSampleAt(t *testing.T) {
	testSourceSetup(44100, 2)
	testFile := "testdata/Float32bitLittleEndian48000HzEstéreo.wav"
	source := New(testFile)
	
	// Test getting sample at beginning
	smp := source.SampleAt(0, 1.0, 0)
	assert.Equal(t, 2, len(smp))
	
	// Test with volume adjustment
	smpHalfVol := source.SampleAt(0, 0.5, 0)
	assert.Equal(t, 2, len(smpHalfVol))
	
	// Test with pan
	smpLeftPan := source.SampleAt(0, 1.0, -1.0)
	assert.Equal(t, 2, len(smpLeftPan))
	
	smpRightPan := source.SampleAt(0, 1.0, 1.0)
	assert.Equal(t, 2, len(smpRightPan))
	
	// Test beyond length should return zeros
	smpBeyond := source.SampleAt(source.Length()+100, 1.0, 0)
	assert.Equal(t, 2, len(smpBeyond))
	assert.Equal(t, sample.Value(0), smpBeyond[0])
	assert.Equal(t, sample.Value(0), smpBeyond[1])
}

func TestState(t *testing.T) {
	testSourceSetup(44100, 1)
	testFile := "testdata/Signed16bitLittleEndian44100HzMono.wav"
	source := New(testFile)
	
	// After loading, state should be READY
	assert.Equal(t, READY, source.state)
}

func TestStateName(t *testing.T) {
	// Test that state enum values are as expected
	assert.Equal(t, stateEnum(0), STAGED)
	assert.Equal(t, stateEnum(1), LOADING)
	assert.Equal(t, stateEnum(2), READY)
}

func TestLength(t *testing.T) {
	testSourceSetup(44100, 1)
	testFile := "testdata/Signed16bitLittleEndian44100HzMono.wav"
	source := New(testFile)
	
	// Length should be greater than 0 for a valid audio file
	length := source.Length()
	assert.True(t, length > 0, "Source length should be greater than 0")
	
	// Length should match maxTz
	assert.Equal(t, source.maxTz, length)
}

func TestTeardown(t *testing.T) {
	testSourceSetup(44100, 1)
	testFile := "testdata/Signed16bitLittleEndian44100HzMono.wav"
	source := New(testFile)
	
	// Source should have samples before teardown
	assert.NotNil(t, source.sample)
	assert.True(t, len(source.sample) > 0)
	
	// Teardown should not panic
	assert.NotPanics(t, func() {
		source.Teardown()
	})
	
	// After teardown, sample should be nil
	assert.Nil(t, source.sample)
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
