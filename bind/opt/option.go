// Package opt specifies valid options
package opt

// OptLoader represents an audio input option
type Input string

// OptLoadWav to use Go-Native WAV file I/O
const InputWAV Input = "wav"

// OptOutput represents an audio output option
type Output string

// OptOutputNull for benchmarking/profiling, because those tools are unable to sample to C-go callback tree
const OutputNull Output = "null"

// OptOutputPortAudio to use Portaudio for audio output
const OutputPortAudio Output = "portaudio"

// OptOutputSDL to use SDL for audio output
const OutputSDL Output = "sdl"

// OptOutputWAV to use WAV directly for []byte to stdout
const OutputWAV Output = "wav"
