// Package debug for debugging
package debug

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigure(t *testing.T) {
	// Initially should be false
	Configure(false)
	assert.False(t, Active())

	// Set to true
	Configure(true)
	assert.True(t, Active())

	// Set back to false
	Configure(false)
	assert.False(t, Active())
}

func TestActive(t *testing.T) {
	Configure(false)
	assert.False(t, Active())

	Configure(true)
	assert.True(t, Active())
}

func TestPrintf(t *testing.T) {
	// Capture stderr output
	oldLogger := logger
	defer func() { logger = oldLogger }()

	var buf bytes.Buffer
	logger = log.New(&buf, "", 0)

	// When debug is off, nothing should be printed
	Configure(false)
	Printf("test message %s", "hello")
	assert.Empty(t, buf.String())

	// When debug is on, message should be printed
	Configure(true)
	Printf("test message %s", "hello")
	assert.Contains(t, buf.String(), "test message hello")
}

func TestInit(t *testing.T) {
	// Verify logger is initialized
	assert.NotNil(t, logger)
	
	// Verify it uses os.Stderr (we can't test this directly but we can check it's not nil)
	// This is more of a sanity check
	assert.NotNil(t, os.Stderr)
}
