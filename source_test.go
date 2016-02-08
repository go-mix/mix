// Package atomix is a sequence-based Go-native audio mixer
package atomix

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/outrightmental/go-atomix/bind"
)

// TODO: test multi-channel source audio files

func TestSource_Base(t *testing.T) {
	// TODO: Test Source Base
}

func TestSource_Load_FAIL(t *testing.T) {
	pathFail := "./lib/ThisShouldFailBecauseItDoesNotExist.wav"
	defer func() {
		msg := recover()
		assert.IsType(t, "", msg)
		assert.Equal(t, "File not found: "+pathFail, msg)
	}()
	Debug(true)
	testSourceSetup(44100, 1)
	source := NewSource(pathFail)
	assert.NotNil(t, source)
}

func TestSource_LoadSigned16bitLittleEndian44100HzMono(t *testing.T) {
	Debug(true)
	testSourceSetup(44100, 1)
	source := NewSource("./lib/Signed16bitLittleEndian44100HzMono.wav")
	assert.NotNil(t, source)
	totalSoundMovement := testSourceAssertSound(t, source, 1)
	assert.True(t, totalSoundMovement > .001)
}

func TestSource_LoadFloat32bitLittleEndian48000HzEstéreo(t *testing.T) {
	Debug(true)
	testSourceSetup(48000, 2)
	source := NewSource("./lib/Float32bitLittleEndian48000HzEstéreo.wav")
	assert.NotNil(t, source)
	totalSoundMovement := testSourceAssertSound(t, source, 2)
	assert.True(t, totalSoundMovement > .001)
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

//
// Test Components
//

func testSourceSetup(freq float64, channels int) {
	Configure(bind.AudioSpec{
		Freq:     freq,
		Format:   bind.AudioF32,
		Channels: channels,
	})
}

func testSourceAssertSound(t *testing.T, source *Source, channels int) (totalSoundMovement float64) {
	for tz := Tz(0); tz < source.Length(); tz++ {
		smp := source.SampleAt(tz, 1, 0)
		assert.Equal(t, channels, len(smp))
		for c := 0; c < channels; c++ {
			totalSoundMovement += math.Abs(smp[c])
		}
	}
	return
}
