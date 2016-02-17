// Package bind is for modular binding of ontomix to audio interface
package bind

import (
	"github.com/gordonklaus/portaudio"
)

var outPortaudioStream *portaudio.Stream

func outPortaudioSetup(spec *AudioSpec) {
	var err error
	portaudio.Initialize()
	outPortaudioStream, err = portaudio.OpenDefaultStream(0, spec.Channels, spec.Freq, 0, outPortaudioStreamCallback)
	noErr(err)
	noErr(outPortaudioStream.Start())
}

func outPortaudioTeardown() {
	//	noErr(out.Stop())
	//	noErr(out.Close())
	portaudio.Terminate()
}

func outPortaudioStreamCallback(out [][]float32) {
	var sample []float64
	for s := range out[0] {
		sample = outCallbackMixNextSample()
		for c := 0; c < outSpec.Channels; c++ {
			out[c][s] = float32(sample[c])
		}
	}
}
