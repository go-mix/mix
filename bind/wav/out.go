package wav

import (
	"gopkg.in/ontomix.v0/bind/spec"
)

var (
	outputSpec *spec.AudioSpec
)

func ConfigureOutput(s spec.AudioSpec) {
	outputSpec = &s
}

func TeardownOutput() {

}
