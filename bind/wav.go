// Package bind has native Go WAV I/O
package bind

import (
	"io"

	"github.com/youpy/go-wav"
	"os"
)


func nativeLoadWAV(path string) (out [][]float64, spec *AudioSpec) {
	//	data, sdlSpec := sdl.LoadWAV(file, sdl2Spec(spec))
	// return data, sdl2Unspec(sdlSpec)
	file, _ := os.Open(path)
	reader := wav.NewReader(file)
	fmt, err := reader.Format()
	if err != nil {
		return
	}
	spec = &AudioSpec{
		Channels: fmt.NumChannels,
	}
	for {
		samples, err := reader.ReadSamples()
//		type WavFormat struct {
//			AudioFormat   uint16
//			NumChannels   uint16
//			SampleRate    uint32
//			ByteRate      uint32
//			BlockAlign    uint16
//			BitsPerSample uint16
//		}
		if err == io.EOF {
			break
		}
		for _, sample := range samples {
			row := make([]float64, 0)
			for c := uint(0); c < uint(fmt.NumChannels); c ++ {
				row = append(row, reader.FloatValue(sample, c))
			}
			out = append(out, row)
		}
	}
	return
}
