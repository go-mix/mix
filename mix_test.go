// Sequence-based Go-native audio mixer for music apps
package mix

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/go-mix/mix/bind/spec"
	"github.com/go-mix/mix/lib/mix"
)

func TestDebug(t *testing.T) {
	// Test turning debug on
	Debug(true)
	// Since debug package is internal, we test indirectly that it doesn't panic
	assert.NotPanics(t, func() {
		Debug(true)
		Debug(false)
	})
}

func TestConfigure(t *testing.T) {
	// Test valid configuration
	assert.NotPanics(t, func() {
		Configure(spec.AudioSpec{
			Freq:     44100,
			Format:   spec.AudioF32,
			Channels: 2,
		})
	})
}

func TestConfigure_FailureFreqNotGreaterThanZero(t *testing.T) {
	defer func() {
		msg := recover()
		assert.IsType(t, "", msg)
		assert.Equal(t, "Must specify a mixing frequency greater than zero.", msg)
	}()
	Configure(spec.AudioSpec{
		Freq:     -100,
		Format:   spec.AudioS16,
		Channels: 2,
	})
}

func TestTeardown(t *testing.T) {
	testAPISetup()
	Teardown()
}

func TestSpec(t *testing.T) {
	testAPISetup()
	assert.Equal(t, &spec.AudioSpec{
		Freq:     44100,
		Format:   spec.AudioF32,
		Channels: 1,
	}, Spec())
	Teardown()
}

func TestSetFire(t *testing.T) {
	testAPISetup()
	fire := SetFire("lib/source/testdata/Signed16bitLittleEndian44100HzMono.wav", time.Duration(0), 0, 1.0, 0)
	assert.NotNil(t, fire)
}

func TestFireCount(t *testing.T) {
	testAPISetup()
	assert.Equal(t, 0, FireCount())
	SetFire("lib/source/testdata/Float32bitLittleEndian48000HzEstéreo.wav", time.Duration(0), 0, 1.0, 0)
	assert.Equal(t, 1, FireCount())
	SetFire("lib/source/testdata/Signed16bitLittleEndian44100HzMono.wav", time.Duration(0), 0, 1.0, 0)
	assert.Equal(t, 2, FireCount())
	// TODO: assert count drains during back to 0 as a result of playback
}

func TestClearAllFires(t *testing.T) {
	testAPISetup()
	SetFire("lib/source/testdata/Signed16bitLittleEndian44100HzMono.wav", time.Duration(0), 0, 1.0, 0)
	ClearAllFires()
	assert.Equal(t, 0, FireCount())
}

func TestSetSoundsPath(t *testing.T) {
	// Test setting sounds path doesn't panic
	assert.NotPanics(t, func() {
		SetSoundsPath("test/path/")
		SetSoundsPath("")
	})
}

func TestSetGetMixCycleDuration(t *testing.T) {
	testAPISetup()
	SetMixCycleDuration(2 * time.Second)
	assert.Equal(t, spec.Tz(88200), mix.GetCycleDurationTz())
}

func TestStart(t *testing.T) {
	Start()
}

func TestStartAt(t *testing.T) {
	StartAt(time.Now().Add(1 * time.Second))
}

func TestGetStartTime(t *testing.T) {
	startExpect := time.Now().Add(1 * time.Second)
	StartAt(startExpect)
	startActual := GetStartTime()
	assert.Equal(t, startExpect, startActual)
}

func TestGetNowAt(t *testing.T) {
	testAPISetup()
	Start()
	// Give it a moment to start
	time.Sleep(10 * time.Millisecond)
	now := GetNowAt()
	// Should be a positive duration after start
	assert.True(t, now >= 0)
}

func TestOutputStart(t *testing.T) {
	testAPISetup()
	// Create a buffer to write to
	var buf bytes.Buffer
	assert.NotPanics(t, func() {
		OutputStart(1*time.Second, &buf)
	})
}

func TestOutputContinueTo(t *testing.T) {
	testAPISetup()
	assert.NotPanics(t, func() {
		OutputContinueTo(100 * time.Millisecond)
	})
}

func TestOutputClose(t *testing.T) {
	testAPISetup()
	assert.NotPanics(t, func() {
		OutputClose()
	})
}

func TestAudioCallback(t *testing.T) {
	// Test that the audio system can be configured
	// This is tested indirectly through Configure
	testAPISetup()
	assert.NotNil(t, Spec())
}

//
// Test Components
//

func testAPISetup() {
	Teardown()
	ClearAllFires()
	Configure(spec.AudioSpec{
		Freq:     44100,
		Format:   spec.AudioF32,
		Channels: 1,
	})
}
