// Package source models a single audio source
package source

import (
	"testing"

	"github.com/go-mix/mix/bind/spec"
	"github.com/stretchr/testify/assert"
)

const (
	testFile1 = "testdata/Signed16bitLittleEndian44100HzMono.wav"
	testFile2 = "testdata/Float32bitLittleEndian48000HzEstéreo.wav"
)

func TestPrepare(t *testing.T) {
	// Setup
	testSourceSetup(44100, 1)

	// Get initial count
	initialCount := Count()

	// Prepare a test source
	Prepare(testFile1)

	// Count should increase by 1
	assert.Equal(t, initialCount+1, Count())

	// Preparing the same source again should not increase count
	Prepare(testFile1)
	assert.Equal(t, initialCount+1, Count())

	// The source should now be retrievable
	src := Get(testFile1)
	assert.NotNil(t, src)
}

func TestGet(t *testing.T) {
	// Setup
	testSourceSetup(44100, 1)

	// Getting a non-existent source should return nil
	src := Get("non-existent-file.wav")
	assert.Nil(t, src)

	// Prepare and get a source
	Prepare(testFile1)
	src = Get(testFile1)
	assert.NotNil(t, src)
	assert.Equal(t, testFile1, src.URL)
}

func TestGetLength(t *testing.T) {
	// Setup
	testSourceSetup(44100, 1)

	// Getting length of non-existent source should return 0
	length := GetLength("non-existent-file.wav")
	assert.Equal(t, spec.Tz(0), length)

	// Prepare a source and get its length
	Prepare(testFile1)
	length = GetLength(testFile1)
	assert.True(t, length > 0, "Length should be greater than 0 for a valid audio file")
}

func TestPrune(t *testing.T) {
	// Setup
	testSourceSetup(44100, 1)

	// Prepare multiple sources
	Prepare(testFile1)
	Prepare(testFile2)

	initialCount := Count()
	assert.True(t, initialCount >= 2, "Should have at least 2 sources")

	// Keep only file1
	keep := map[string]bool{
		testFile1: true,
	}
	Prune(keep)

	// file1 should still exist
	assert.NotNil(t, Get(testFile1))

	// file2 should be pruned
	assert.Nil(t, Get(testFile2))
}

func TestCount(t *testing.T) {
	// Setup
	testSourceSetup(44100, 1)

	// Clear storage by pruning everything
	Prune(map[string]bool{})

	// Count should be 0 or very low
	count := Count()

	// Prepare sources and verify count increases
	Prepare(testFile1)
	assert.Equal(t, count+1, Count())

	Prepare(testFile2)
	assert.Equal(t, count+2, Count())
}
