// Package source models a single audio source
package source

import (
	"sync"
)

// Prepare a source by ensuring it is stored in memory.
func Prepare(src string) {
	storageMutex.Lock()
	defer storageMutex.Unlock()
	if _, exists := storage[src]; !exists {
		storage[src] = New(src)
	}
}

// Get a source from storage
func Get(src string) *Source {
	storageMutex.Lock()
	defer storageMutex.Unlock()
	if _, ok := storage[src]; ok {
		return storage[src]
	} else {
		return nil
	}
}

// Prune to keep only the sources in this list
func Prune(keep map[string]bool) {
	storageMutex.Lock()
	defer storageMutex.Unlock()
	for key, _ := range storage {
		if _, exists := keep[key]; !exists {
			delete(storage, key)
		}
	}
}

// Count the number of sources in memory
func Count() int {
	storageMutex.Lock()
	defer storageMutex.Unlock()
	return len(storage)
}

/*
 *
 private */

var (
	storage      map[string]*Source
	storageMutex = &sync.Mutex{}
)

func init() {
	storage = make(map[string]*Source, 0)
}
