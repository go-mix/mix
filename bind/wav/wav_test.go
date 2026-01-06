// Package wav is direct WAV filo I/O
package wav

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	// TODO
}

func TestLoadFromReader(t *testing.T) {
	// Read a test WAV file
	wavData, err := ioutil.ReadFile("../../lib/source/testdata/Signed16bitLittleEndian44100HzMono.wav")
	assert.NoError(t, err)
	
	// Create a bytes.Reader which implements both io.Reader and io.ReaderAt
	reader := bytes.NewReader(wavData)
	
	// Load from reader
	samples, spec := LoadFromReader(reader)
	
	// Verify we got samples and spec
	assert.NotNil(t, samples)
	assert.NotNil(t, spec)
	assert.Greater(t, len(samples), 0)
	assert.Greater(t, spec.Freq, float64(0))
	assert.Greater(t, spec.Channels, 0)
}
