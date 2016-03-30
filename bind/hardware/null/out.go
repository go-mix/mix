// Package null is for modular binding of mix to a null (mock) audio interface
package null

import (
	"github.com/go-mix/mix/bind/sample"
	"github.com/go-mix/mix/bind/spec"
)

func ConfigureOutput(s spec.AudioSpec) {
	go func() {
		for {
			sample.OutNextBytes()
		}
	}()
	// nothing to do
}
