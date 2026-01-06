// Package wav is direct WAV filo I/O
package wav

import (
	"io"
	"os"

	riff "github.com/youpy/go-riff"

	"github.com/go-mix/mix/bind/sample"
	"github.com/go-mix/mix/bind/spec"
)

// Load a WAV file into memory
func Load(path string) (out []sample.Sample, specs *spec.AudioSpec) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("File not found: " + path)
	}
	file, _ := os.Open(path)
	return LoadFromReader(file)
}

// LoadFromReader loads WAV data from an io.Reader+io.ReaderAt into memory
func LoadFromReader(r riff.RIFFReader) (out []sample.Sample, specs *spec.AudioSpec) {
	reader, err := NewReader(r)
	if err != nil {
		panic(err)
	}
	specs = &spec.AudioSpec{
		Freq:     float64(reader.Format.SampleRate),
		Format:   reader.AudioFormat,
		Channels: int(reader.Format.NumChannels),
	}
	for {
		samples, err := reader.ReadSamples()
		if err == io.EOF {
			break
		}
		out = append(out, samples...)
	}
	return
}

//
// Private
//
