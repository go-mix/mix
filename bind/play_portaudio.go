// Package bind is for modular binding of atomix to audio interface
package bind

import (
	"github.com/gordonklaus/portaudio"
)

var playPortaudioStream *portaudio.Stream

func playPortaudioSetup(spec *AudioSpec) {
	var err error
	portaudio.Initialize()
	playPortaudioStream, err = portaudio.OpenDefaultStream(0, spec.Channels, spec.Freq, 0, playPortaudioStreamCallback)
	noErr(err)
	noErr(playPortaudioStream.Start())
}

func playPortaudioTeardown() {
	//	noErr(output.Stop())
	//	noErr(output.Close())
	portaudio.Terminate()
}

func playPortaudioStreamCallback(out [][]float32) {
	var sample []float64
	for s := range out[0] {
		sample = outputCallbackMixNextSample()
		for c := 0; c < outputSpec.Channels; c++ {
			out[c][s] = float32(sample[c])
		}
	}
}
