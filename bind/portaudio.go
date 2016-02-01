// Package bind is for modular binding of atomix to audio interface
package bind

import (
	"github.com/gordonklaus/portaudio"
)

type portaudioOutput struct {
	*portaudio.Stream
	spec *AudioSpec
}

var output portaudioOutput

func portaudioSetup(spec *AudioSpec) {
	var err error
	portaudio.Initialize()
	output := portaudioOutput{spec: spec}
	output.Stream, err = portaudio.OpenDefaultStream(0, spec.Channels, spec.Freq, 0, output.processAudio)
	noErr(err)
	noErr(output.Start())
}

func portaudioTeardown() {
//	noErr(output.Stop())
//	noErr(output.Close())
	portaudio.Terminate()
}

func (o *portaudioOutput) processAudio(out [][]float32) {
	var sample []float64
	for s := range out[0] {
		sample = mixNextOutputSample()
		for c := 0; c < o.spec.Channels; c++ {
			out[c][s] = float32(sample[c])
		}
	}
}
