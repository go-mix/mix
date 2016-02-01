// Package atomix is a sequence-based Go-native audio mixer
package atomix

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/outrightmental/go-atomix/bind"
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

//
// Test Components
//

func testSourceSetup() {
	Configure(bind.AudioSpec{
		Freq:     44100,
		Format:   bind.AudioF32,
		Channels: 1,
	})
}
