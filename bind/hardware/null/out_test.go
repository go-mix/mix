// Package null is for modular binding of mix to a null (mock) audio interface
package null

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/go-mix/mix/bind/spec"
)

func TestConfigureOutput(t *testing.T) {
	// Test that ConfigureOutput doesn't panic
	s := spec.AudioSpec{
		Freq:     44100,
		Format:   spec.AudioF32,
		Channels: 2,
	}
	
	assert.NotPanics(t, func() {
		ConfigureOutput(s)
	})
}
