package wav

import (
	"github.com/go-ontomix/ontomix/bind/spec"
)

var (
	outputSpec *spec.AudioSpec
)

func ConfigureOutput(s spec.AudioSpec) {
	outputSpec = &s
}

func TeardownOutput() {

}
