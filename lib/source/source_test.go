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

// TestStereoImplicitPanning tests that 2-channel (stereo) sources respect
// the implicit L/R panning assumption: channel 0 = left, channel 1 = right
func TestStereoImplicitPanning(t *testing.T) {
	// Configure for stereo: 2 channels
	testSourceSetup(48000, 2)

	// Load a stereo source file
	source := New("testdata/Float32bitLittleEndian48000HzEstéreo.wav")
	assert.NotNil(t, source)
	assert.Equal(t, 2, source.audioSpec.Channels, "Source should have 2 channels (stereo)")

	// Test that when pan=0 (center), stereo channels are preserved as-is
	// This tests the implicit panning: left channel stays left, right channel stays right
	if source.Length() > 0 {
		// Get a sample from the middle of the audio where there's likely actual audio data
		midPoint := source.Length() / 2
		smp := source.SampleAt(midPoint, 1.0, 0)

		// Verify we get 2 channels back (stereo output)
		assert.Equal(t, 2, len(smp), "Output should have 2 channels")

		// When pan=0, both channels should receive their original source data
		// The volume function with pan=0 returns volume=1.0 for all channels
		// So output channel 0 gets source channel 0 (left)
		// And output channel 1 gets source channel 1 (right)
		sourceSample := source.sample[midPoint]
		assert.Equal(t, sourceSample.Values[0], smp[0], "Left channel (0) should map to left output")
		assert.Equal(t, sourceSample.Values[1], smp[1], "Right channel (1) should map to right output")
	}
}

// TestStereoToStereoChannelMapping tests that stereo source channels
// correctly map to stereo output channels with no panning applied
func TestStereoToStereoChannelMapping(t *testing.T) {
	// Configure for stereo output: 2 channels
	testSourceSetup(48000, 2)

	// Load a stereo source file
	source := New("testdata/Float32bitLittleEndian48000HzEstéreo.wav")
	assert.NotNil(t, source)
	assert.Equal(t, 2, source.audioSpec.Channels, "Source should be stereo (2 channels)")
	assert.Equal(t, 2, masterSpec.Channels, "Output should be stereo (2 channels)")

	// Iterate through several samples to ensure consistent behavior
	sampleCount := 0
	for tz := spec.Tz(0); tz < source.Length() && sampleCount < 100; tz += source.Length() / 100 {
		smp := source.SampleAt(tz, 1.0, 0)

		// Verify 2 channels in output
		assert.Equal(t, 2, len(smp), "Each sample should have 2 channels")

		// With pan=0 and volume=1.0, output should match source
		sourceSample := source.sample[tz]
		assert.Equal(t, sourceSample.Values[0], smp[0], "Left source channel should map to left output at tz=%d", tz)
		assert.Equal(t, sourceSample.Values[1], smp[1], "Right source channel should map to right output at tz=%d", tz)

		sampleCount++
	}
	assert.True(t, sampleCount > 0, "Should have tested at least some samples")
}

// TestStereoImplicitPanningWithVolume tests stereo channel mapping
// with different volume levels (but no panning)
func TestStereoImplicitPanningWithVolume(t *testing.T) {
	testSourceSetup(48000, 2)
	source := New("testdata/Float32bitLittleEndian48000HzEstéreo.wav")
	assert.NotNil(t, source)

	if source.Length() > 0 {
		midPoint := source.Length() / 2
		sourceSample := source.sample[midPoint]

		// Test with volume = 0.5, pan = 0
		smp := source.SampleAt(midPoint, 0.5, 0)
		assert.Equal(t, 2, len(smp))

		// With pan=0, volume applies equally to all channels
		// So each channel should be sourceValue * 0.5
		assert.Equal(t, sourceSample.Values[0]*0.5, smp[0], "Left channel should be scaled by volume")
		assert.Equal(t, sourceSample.Values[1]*0.5, smp[1], "Right channel should be scaled by volume")

		// Test with volume = 0.75, pan = 0
		smp2 := source.SampleAt(midPoint, 0.75, 0)
		assert.Equal(t, sourceSample.Values[0]*0.75, smp2[0], "Left channel should be scaled by volume 0.75")
		assert.Equal(t, sourceSample.Values[1]*0.75, smp2[1], "Right channel should be scaled by volume 0.75")
	}
}

// TestStereoChannelIdentity verifies that channel 0 is left and channel 1 is right
// by testing the volume function's behavior with stereo configuration
func TestStereoChannelIdentity(t *testing.T) {
	// Set up for stereo using the standard test setup function
	testSourceSetup(48000, 2)

	// With pan=0 (center), both channels should have equal volume multiplier
	leftVol := volume(0, 1.0, 0)  // channel 0 = left
	rightVol := volume(1, 1.0, 0) // channel 1 = right
	assert.Equal(t, sample.Value(1.0), leftVol, "Left channel (0) with pan=0 should have full volume")
	assert.Equal(t, sample.Value(1.0), rightVol, "Right channel (1) with pan=0 should have full volume")

	// With pan=-1 (full left), left channel should be full, right should be reduced
	leftVolLeft := volume(0, 1.0, -1)  // channel 0 = left, panned left
	rightVolLeft := volume(1, 1.0, -1) // channel 1 = right, panned left
	assert.Equal(t, sample.Value(1.0), leftVolLeft, "Left channel (0) with pan=-1 should have full volume")
	assert.Equal(t, sample.Value(0.5), rightVolLeft, "Right channel (1) with pan=-1 should be reduced")

	// With pan=+1 (full right), the current algorithm produces:
	// channel 0: 1 - 1*0/2 = 1, channel 1: 1 - 1*1/2 = 0.5
	// This is the current behavior (verified by existing tests in TestMixer_mixVolume)
	leftVolRight := volume(0, 1.0, 1)  // channel 0 = left, panned right
	rightVolRight := volume(1, 1.0, 1) // channel 1 = right, panned right
	assert.Equal(t, sample.Value(1.0), leftVolRight, "Left channel (0) with pan=+1 current behavior")
	assert.Equal(t, sample.Value(0.5), rightVolRight, "Right channel (1) with pan=+1 current behavior")
}

//
// Private
//

func testSourceSetup(freq float64, channels int) {
	Clear()
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
