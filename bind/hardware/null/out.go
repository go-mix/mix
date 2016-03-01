// Package null is for modular binding of ontomix to a null (mock) audio interface
package null

import (
	"gopkg.in/ontomix.v0/bind/sample"
	"gopkg.in/ontomix.v0/bind/spec"
)

func ConfigureOutput(s spec.AudioSpec) {
	go func() {
		for {
			sample.OutNextBytes()
		}
	}()
	// nothing to do
}
