// Package null is for modular binding of ontomix to a null (mock) audio interface
package null

import (
	"github.com/go-ontomix/ontomix/bind/sample"
	"github.com/go-ontomix/ontomix/bind/spec"
)

func ConfigureOutput(s spec.AudioSpec) {
	go func() {
		for {
			sample.OutNextBytes()
		}
	}()
	// nothing to do
}
