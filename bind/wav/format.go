// Package wav is direct WAV filo I/O
package wav

import (
	"github.com/go-ontomix/ontomix/bind/spec"
	"io"
)

// the Format struct must be in the exact order according
// to WAV specifications, such that a binary.Read(...)
// can assign the WAV specified "fmt" header bytes
// to the correct Format properties.
type Format struct {
	SampleFormat  SampleFormat
	NumChannels   uint16
	SampleRate    uint32
	ByteRate      uint32
	BlockAlign    uint16
	BitsPerSample uint16
}

func FormatFromSpec(s *spec.AudioSpec) Format {
	format := Format{}
	switch s.Format {
	case spec.AudioU8:
		format.SampleFormat = AudioFormatLinearPCM
		format.BitsPerSample = 8
	case spec.AudioS8:
		format.SampleFormat = AudioFormatLinearPCM
		format.BitsPerSample = 8
	case spec.AudioU16:
		format.SampleFormat = AudioFormatLinearPCM
		format.BitsPerSample = 16
	case spec.AudioS16:
		format.SampleFormat = AudioFormatLinearPCM
		format.BitsPerSample = 16
	case spec.AudioS32:
		format.SampleFormat = AudioFormatLinearPCM
		format.BitsPerSample = 32
	case spec.AudioF32:
		format.SampleFormat = AudioFormatIEEEFloat
		format.BitsPerSample = 32
	case spec.AudioF64:
		format.SampleFormat = AudioFormatIEEEFloat
		format.BitsPerSample = 64
	}
	format.NumChannels = uint16(s.Channels)
	format.SampleRate = uint32(s.Freq)
	if format.ByteRate == 0 {
		format.ByteRate = format.SampleRate * uint32(format.NumChannels*format.BitsPerSample/8)
	}
	if format.BlockAlign == 0 {
		format.BlockAlign = format.NumChannels * format.BitsPerSample / 8
	}
	return format
}

type Data struct {
	io.Reader
	Size uint32
	pos  uint32
}

type SampleFormat uint16

const (
	AudioFormatLinearPCM SampleFormat = 0x0001
	AudioFormatIEEEFloat SampleFormat = 0x0003
)
