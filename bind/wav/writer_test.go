// Package wav is direct WAV filo I/O
package wav

import (
	//"io/ioutil"
	//"os"
	//"log"
	"testing"
	//"github.com/stretchr/testify/assert"
)

func TestConfigureOutput(t *testing.T) {
	// TODO
}

func TestTeardownOutput(t *testing.T) {
	// TODO
}

func TestWrite(t *testing.T) {
	//outfile, err := ioutil.TempFile("/tmp", "outfile")
	//if err != nil {
	//	t.Fatal(err)
	//}
	//log.Printf("TempFile: %+v", outfile.Name())
	//
	//defer func() {
	//	outfile.Close()
	//	os.Remove(outfile.Name())
	//}()
	//
	//var numSamples uint32 = 2
	//var numChannels uint16 = 2
	//var sampleRate uint32 = 44100
	//var bitsPerSample uint16 = 32
	//
	//writer := NewWriter(outfile, AudioFormatIEEEFloat, numSamples, numChannels, sampleRate, bitsPerSample)
	//if writer == nil {
	//	t.Fatal(err)
	//}
	//
	//outfile.Close()
	//file, err := os.Open(outfile.Name())
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//defer func() {
	//	file.Close()
	//	os.Remove(outfile.Name())
	//}()
	//
	//reader, err := NewReader(file)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//assert.Equal(t, reader.Format.SampleFormat, AudioFormatIEEEFloat)
	//assert.Equal(t, reader.Format.NumChannels, numChannels)
	//assert.Equal(t, reader.Format.SampleRate, sampleRate)
	//assert.Equal(t, reader.Format.ByteRate, sampleRate * 8)
	//assert.Equal(t, reader.Format.BlockAlign, numChannels*(bitsPerSample/8))
	//assert.Equal(t, reader.Format.BitsPerSample, bitsPerSample)

}
