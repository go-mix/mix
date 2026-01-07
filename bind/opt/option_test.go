// Package opt specifies valid options
package opt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInput_Constants(t *testing.T) {
	// Test Input type constants
	assert.Equal(t, Input("wav"), InputWAV)
	assert.Equal(t, Input("sox"), InputSOX)
}

func TestInput_TypeString(t *testing.T) {
	// Test that Input type can be converted to string
	wav := InputWAV
	assert.Equal(t, "wav", string(wav))

	sox := InputSOX
	assert.Equal(t, "sox", string(sox))
}

func TestOutput_Constants(t *testing.T) {
	// Test Output type constants
	assert.Equal(t, Output("null"), OutputNull)
	assert.Equal(t, Output("wav"), OutputWAV)
}

func TestOutput_TypeString(t *testing.T) {
	// Test that Output type can be converted to string
	null := OutputNull
	assert.Equal(t, "null", string(null))

	wav := OutputWAV
	assert.Equal(t, "wav", string(wav))
}
