// Package wav is direct WAV filo I/O
package wav

import (
	"bufio"
	"encoding/binary"
	"errors"

	riff "github.com/youpy/go-riff"

	"fmt"
	"gopkg.in/mix.v0/bind/sample"
	"gopkg.in/mix.v0/bind/spec"
)

type Reader struct {
	Format      *Format
	AudioFormat spec.AudioFormat
	*Data
	// private
	riffReader *riff.Reader
	riffChunk  *riff.RIFFChunk
}

func NewReader(file riff.RIFFReader) (reader *Reader, err error) {
	reader = &Reader{riffReader: riff.NewReader(file)}
	format, audioFormat, err := reader.openAndParse()
	if err != nil {
		return
	}
	reader.Format = format
	reader.AudioFormat = audioFormat
	return
}

func (r *Reader) ReadSamples(params ...uint32) (out []sample.Sample, err error) {
	var buffer []byte
	var numSamples, n int

	if len(params) > 0 {
		numSamples = int(params[0])
	} else {
		numSamples = 2048
	}

	numChannels := int(r.Format.NumChannels)
	bytesPerSample := int(r.Format.BitsPerSample) / 8
	blockAlign := int(r.Format.NumChannels) * bytesPerSample

	buffer = make([]byte, numSamples*blockAlign)
	n, err = r.readSamplesIntoBuffer(buffer)

	if err != nil {
		return
	}

	numSamples = n / blockAlign
	r.Data.pos += uint32(numSamples * blockAlign)

	for offset := 0; offset < len(buffer)-numChannels-bytesPerSample; offset += blockAlign {
		values := make([]sample.Value, numChannels)
		for c := 0; c < int(numChannels); c++ {
			offsetCh := offset + c*bytesPerSample
			bytes := buffer[offsetCh : offsetCh+bytesPerSample]
			values[c] = r.sampleFromBytes(r.AudioFormat, bytes)
		}
		out = append(out, sample.New(values))
	}

	return
}

//
// Private
//

func (r *Reader) readSamplesIntoBuffer(p []byte) (n int, err error) {
	if r.Data == nil {
		data, err := r.readData()
		if err != nil {
			return n, err
		}
		r.Data = data
	}

	return r.Data.Read(p)
}

func (r *Reader) sampleFromBytes(audio spec.AudioFormat, bytes []byte) sample.Value {
	// TODO: big-endian or little-endian?
	switch audio {
	case spec.AudioU8:
		return sample.ValueOfByteU8(bytes[0])
	case spec.AudioS8:
		return sample.ValueOfByteS8(bytes[0])
	case spec.AudioU16:
		return sample.ValueOfBytesU16LSB(bytes)
	case spec.AudioS16:
		return sample.ValueOfBytesS16LSB(bytes)
	case spec.AudioS32:
		return sample.ValueOfBytesS32LSB(bytes)
	case spec.AudioF32:
		return sample.ValueOfBytesF32LSB(bytes)
	case spec.AudioF64:
		return sample.ValueOfBytesF64LSB(bytes)
	default:
		panic("Unhandled format!")
	}
}

func (r *Reader) openAndParse() (format *Format, audio spec.AudioFormat, err error) {
	var riffChunk *riff.RIFFChunk

	format = new(Format)

	if r.riffChunk == nil {
		riffChunk, err = r.riffReader.Read()
		if err != nil {
			return
		}

		r.riffChunk = riffChunk
	} else {
		riffChunk = r.riffChunk
	}

	for _, ch := range riffChunk.Chunks {
		var data []byte
		switch string(ch.ChunkID[:]) {
		case "fmt ":
			err = binary.Read(ch, binary.LittleEndian, format)
			if err != nil {
				return
			}
			switch SampleFormat(format.SampleFormat) {
			case AudioFormatLinearPCM: // Linear PCM
				switch format.BitsPerSample {
				case 8:
					audio = spec.AudioS8
				case 16:
					audio = spec.AudioS16
				default:
					panic(fmt.Sprintf("Unhandled Linear PCM bitrate: %+v", format.BitsPerSample))
				}
			case AudioFormatIEEEFloat: // IEEE Float
				switch format.BitsPerSample {
				case 32:
					audio = spec.AudioF32
				case 64:
					audio = spec.AudioF64
				default:
					panic(fmt.Sprintf("Unhandled IEEE Float bitrate: %+v", format.BitsPerSample))
				}
			default:
				panic("Unhandled format")
			}
		case "fact":
			data = make([]byte, ch.ChunkSize)
			err = binary.Read(ch, binary.LittleEndian, data)
			if err != nil {
				return
			}
		case "PEAK":
			data = make([]byte, ch.ChunkSize)
			err = binary.Read(ch, binary.LittleEndian, data)
			if err != nil {
				return
			}
		}
	}

	if format == nil && err == nil {
		err = errors.New("Format chunk is not found")
	}
	return
}

func (r *Reader) readData() (data *Data, err error) {
	var riffChunk *riff.RIFFChunk

	if r.riffChunk == nil {
		riffChunk, err = r.riffReader.Read()
		if err != nil {
			return
		}

		r.riffChunk = riffChunk
	} else {
		riffChunk = r.riffChunk
	}

	for _, ch := range riffChunk.Chunks {
		if string(ch.ChunkID[:]) == "data" {
			data = &Data{bufio.NewReader(ch), ch.ChunkSize, 0}
			return
		}
	}

	err = errors.New("Data chunk is not found")
	return
}
