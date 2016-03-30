// Package wav is direct WAV filo I/O
package wav

import (
	"encoding/binary"
	"io"
	"syscall"

	riff "github.com/youpy/go-riff"

	"gopkg.in/mix.v0/bind/sample"
	"gopkg.in/mix.v0/bind/spec"
	"os"
	"time"
)

func ConfigureOutput(s spec.AudioSpec) {
	outputSpec = &s
}

func OutputStart(length time.Duration) {
	writer = NewWriter(stdout, FormatFromSpec(outputSpec), length)
}

func TeardownOutput() {
	// nothing to do
}

type Writer struct {
	io.Writer
	Format *Format
}

func NewWriter(w io.Writer, format Format, length time.Duration) (writer *Writer) {
	dataSize := uint32(float64(length/time.Second)*float64(format.SampleRate)) * uint32(format.BlockAlign)
	riffSize := 4 + 8 + 16 + 8 + dataSize
	riffWriter := riff.NewWriter(w, []byte("WAVE"), riffSize)

	writer = &Writer{riffWriter, &format}
	riffWriter.WriteChunk([]byte("fmt "), 16, func(w io.Writer) {
		binary.Write(w, binary.LittleEndian, format)
	})
	riffWriter.WriteChunk([]byte("data"), dataSize, func(w io.Writer) {})

	return writer
}

func OutputNext(numSamples spec.Tz) (err error) {
	for n := spec.Tz(0); n < numSamples; n++ {
		writer.Write(sample.OutNextBytes())
	}
	return
}

/*
 *
 private */

var (
	stdout     = os.NewFile(uintptr(syscall.Stdout), "/dev/stdout")
	writer     *Writer
	outputSpec *spec.AudioSpec
)
