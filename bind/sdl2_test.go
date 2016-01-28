// Package bind is for modular binding of atomix to audio interface
package bind

import (
	"testing"
	"unsafe"
)

func TestSDL2(t *testing.T) {
	// TODO
}

func TestSDL2_AudioCallback(t *testing.T) {
	AudioCallback(unsafe.Pointer{}, &uint8(5), int(25))
}
