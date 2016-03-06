// Package debug for debugging
package debug

import (
	"log"
	"os"
)

// Debug ON/OFF (ripples down to all sub-modules)
func Configure(active bool) {
	isActive = active
}

// Printf only when debug is active
func Printf(format string, args ...interface{}) {
	if isActive {
		logger.Printf(format, args...)
	}
}

// Active returns current state of debug
func Active() bool {
	return isActive
}

/*
 *
 private */

var (
	isActive bool
	logger *log.Logger
)

func init() {
	logger = log.New(os.Stderr, "", 0)
}
