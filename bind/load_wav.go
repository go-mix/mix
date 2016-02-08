// Package bind has native Go WAV I/O
package bind

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
	"os"

	riff "github.com/youpy/go-riff"
)

func LoadNewWAV(path string) (out [][]float64, spec *AudioSpec) {
	//	data, sdlSpec := sdl.LoadWAV(file, sdl2Spec(spec))
	// return data, sdl2Unspec(sdlSpec)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("File not found: " + path)
	}
	file, _ := os.Open(path)
	riffReader := riff.NewReader(file)
	reader := &loadWAV{riffReader: riffReader}
	format, audio, err := reader.Open()
	if err != nil {
		return
	}
	spec = &AudioSpec{
		Freq:     float64(format.SampleRate),
		Format:   audio,
		Channels: int(format.NumChannels),
	}
	for {
		samples, err := reader.ReadSamples()
		if err == io.EOF {
			break
		}
		for _, sample := range samples {
			out = append(out, sample)
		}
	}
	return
}

/*
 *
 private */

type loadWAV struct {
	riffReader *riff.Reader
	riffChunk  *riff.RIFFChunk
	format     *loadWAVFormat
	audio      AudioFormat
	*loadWAVData
}

func (r *loadWAV) Open() (format *loadWAVFormat, audio AudioFormat, err error) {
	if r.format == nil {
		format, audio, err = r.openAndParse()
		if err != nil {
			return
		}
		r.format = format
		r.audio = audio
	} else {
		format = r.format
		audio = r.audio
	}

	return
}

func (r *loadWAV) ReadSamples(params ...uint32) (out [][]float64, err error) {
	var buffer []byte
	var numSamples, n int

	if len(params) > 0 {
		numSamples = int(params[0])
	} else {
		numSamples = 2048
	}

	format, audio, err := r.Open()
	if err != nil {
		return
	}

	numChannels := int(format.NumChannels)
	blockAlign := int(format.BlockAlign)
	bytesPerSample := int(format.BitsPerSample) / 8

	buffer = make([]byte, numSamples*blockAlign)
	n, err = r.readSamplesIntoBuffer(buffer)

	if err != nil {
		return
	}

	numSamples = n / blockAlign
	r.loadWAVData.pos += uint32(numSamples * blockAlign)

	for offset := 0; offset < len(buffer)-numChannels-bytesPerSample; offset += blockAlign {
		row := make([]float64, numChannels)
		for c := 0; c < int(numChannels); c++ {
			offsetCh := offset + c*bytesPerSample
			bytes := buffer[offsetCh : offsetCh+bytesPerSample]
			row[c] = r.sampleFromBytes(audio, bytes)
		}
		//		fmt.Printf("append out, %v\n", row)
		out = append(out, row)
	}

	return
}

func (r *loadWAV) readSamplesIntoBuffer(p []byte) (n int, err error) {
	if r.loadWAVData == nil {
		data, err := r.readData()
		if err != nil {
			return n, err
		}
		r.loadWAVData = data
	}

	return r.loadWAVData.Read(p)
}

func (r *loadWAV) sampleFromBytes(audio AudioFormat, bytes []byte) float64 {
	// TODO: big-endian or little-endian?
	switch audio {
	case AudioU8:
		return sampleByteU8(bytes[0])
	case AudioS8:
		return sampleByteS8(bytes[0])
	case AudioU16:
		return sampleBytesU16LSB(bytes)
	case AudioS16:
		return sampleBytesS16LSB(bytes)
	case AudioS32:
		return sampleBytesS32LSB(bytes)
	case AudioF32:
		return sampleBytesF32LSB(bytes)
	case AudioF64:
		return sampleBytesF64LSB(bytes)
	default:
		panic("Unhandled format!")
	}
}

func (r *loadWAV) openAndParse() (format *loadWAVFormat, audio AudioFormat, err error) {
	var riffChunk *riff.RIFFChunk

	format = new(loadWAVFormat)

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
			switch loadWAVAudio(format.AudioFormat) {
			case loadWAVAudioLinearPCM: // Linear PCM
				switch format.BitsPerSample {
				case 8:
					audio = AudioS8
				case 16:
					audio = AudioS16
				default:
					panic("Unhandled Linear PCM bitrate")
				}
			case loadWAVAudioIEEEFloat: // IEEE Float
				switch format.BitsPerSample {
				case 32:
					audio = AudioF32
				case 64:
					audio = AudioF64
				default:
					panic("Unhandled IEEE Float bitrate")
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

func (r *loadWAV) readData() (data *loadWAVData, err error) {
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
			data = &loadWAVData{bufio.NewReader(ch), ch.ChunkSize, 0}
			return
		}
	}

	err = errors.New("Data chunk is not found")
	return
}

type loadWAVFormat struct {
	AudioFormat   uint16
	NumChannels   uint16
	SampleRate    uint32
	ByteRate      uint32
	BlockAlign    uint16
	BitsPerSample uint16
}

type loadWAVData struct {
	io.Reader
	Size uint32
	pos  uint32
}

type loadWAVAudio uint16

const (
	loadWAVAudioLinearPCM loadWAVAudio = 0x0001
	loadWAVAudioIEEEFloat loadWAVAudio = 0x0003
)
