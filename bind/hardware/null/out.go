// Package null is for modular binding of mix to a null (mock) audio interface
package null

import (
	"gopkg.in/mix.v0/bind/sample"
	"gopkg.in/mix.v0/bind/spec"
)

func ConfigureOutput(s spec.AudioSpec) {
	go func() {
		for {
			sample.OutNextBytes()
		}
	}()
	// nothing to do
}
