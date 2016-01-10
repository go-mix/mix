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
