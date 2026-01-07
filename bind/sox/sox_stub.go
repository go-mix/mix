// +build !sox

// Package sox is for file I/O via go-sox package (stub implementation)
package sox

import (
	"github.com/go-mix/mix/bind/sample"
	"github.com/go-mix/mix/bind/spec"
)

const ChunkSize = 2048

// Load sound file into memory (stub - panics when called)
func Load(path string) (out []sample.Sample, specs *spec.AudioSpec) {
	panic("Sox support not compiled in. Use WAV loader instead.")
}
