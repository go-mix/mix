// Package bind is for modular binding of atomix to audio interface
package bind

import (
	"github.com/gordonklaus/portaudio"
)

var portaudioStream *portaudio.Stream

func portaudioSetup(spec *AudioSpec) {
	var err error
	portaudio.Initialize()
	portaudioStream, err = portaudio.OpenDefaultStream(0, spec.Channels, spec.Freq, 0, portaudioStreamCallback)
	noErr(err)
	noErr(portaudioStream.Start())
}

func portaudioTeardown() {
	//	noErr(output.Stop())
	//	noErr(output.Close())
	portaudio.Terminate()
}

func portaudioStreamCallback(out [][]float32) {
	var sample []float64
	for s := range out[0] {
		sample = outputCallbackSample()
		for c := 0; c < outputSpec.Channels; c++ {
			out[c][s] = float32(sample[c])
		}
	}
}
