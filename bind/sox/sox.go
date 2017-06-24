// Package sox is for file I/O via go-sox package
package sox

import (
	"github.com/go-mix/mix/bind/sample"
	"github.com/go-mix/mix/bind/spec"
	sox "github.com/krig/go-sox"
)

const ChunkSize = 2048

// Load sound file into memory
func Load(path string) (out []sample.Sample, specs *spec.AudioSpec) {
	file := sox.OpenRead(path)
	if file == nil {
		panic("Sox can't open file: " + path)
	}
	defer file.Release()
	info := file.Signal()
	specs = &spec.AudioSpec{
		Freq:     info.Rate(),
		Format:   spec.AudioS32,
		Channels: int(info.Channels()),
	}
	buffer := make([]sox.Sample, ChunkSize*specs.Channels)
	for {
		size := file.Read(buffer, uint(len(buffer)))
		if size == 0 || size == sox.EOF {
			break
		}
		outBuffer := make([]sample.Value, size)
		for offset := 0; offset < int(size); offset += specs.Channels {
			values := outBuffer[offset : offset+specs.Channels]
			for c := 0; c < specs.Channels; c++ {
				values[c] = sample.Value(sox.SampleToFloat64(buffer[offset+c]))
			}
			out = append(out, sample.New(values))
		}
	}
	return
}
