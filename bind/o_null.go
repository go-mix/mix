// Package bind is for modular binding of ontomix to audio interface
package bind

import ()

func outNullSetup(spec *AudioSpec) {
	go func() {
		for {
			outNextSample()
		}
	}()
	// nothing to do
}
