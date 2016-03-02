// Package wav is direct WAV filo I/O
package wav

import (
	"encoding/binary"
	"io"
	"syscall"

	riff "github.com/youpy/go-riff"

	"github.com/go-ontomix/ontomix/bind/spec"
	"github.com/go-ontomix/ontomix/bind/sample"
	"os"
	"time"
)

func ConfigureOutput(s spec.AudioSpec) {
	outputSpec = &s
	writer = NewWriter(stdout, 1 * time.Minute, FormatFromSpec(outputSpec))
	// TODO: create a new writer to stdout
}

func TeardownOutput() {
	// nothing to do
}

type Writer struct {
	io.Writer
	Format *Format
}

func NewWriter(w io.Writer, length time.Duration, format Format) (writer *Writer) {
	dataSize := uint32(float64(length / time.Second) * float64(format.SampleRate)) * uint32(format.BlockAlign)
	riffSize := 4 + 8 + 16 + 8 + dataSize
	riffWriter := riff.NewWriter(w, []byte("WAVE"), riffSize)

	writer = &Writer{riffWriter, &format}
	riffWriter.WriteChunk([]byte("fmt "), 16, func(w io.Writer) {
		binary.Write(w, binary.LittleEndian, format)
	})
	riffWriter.WriteChunk([]byte("data"), dataSize, func(w io.Writer) {})

	return writer
}

func WriteSamples(numSamples spec.Tz) (err error) {
	for n := spec.Tz(0); n < numSamples; n++ {
		writer.Write(sample.OutNextBytes())
	}
	return
}

/*
 *
 private */

var (
	stdout = os.NewFile(uintptr(syscall.Stdout), "/dev/stdout")
	writer *Writer
	outputSpec *spec.AudioSpec
)
