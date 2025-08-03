package data

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

// cacheEntry holds the result of a single lazy-loaded JSON file.
type cacheEntry struct {
	once sync.Once
	data any
	err  error
}

var cache = make(map[string]*cacheEntry)
var cacheLock sync.Mutex

// LoadJSON lazily loads and unmarshals a JSON file into a slice of T.
func LoadJSON[T any](path string) ([]T, error) {
	cacheLock.Lock()
	entry, exists := cache[path]
	if !exists {
		entry = &cacheEntry{}
		cache[path] = entry
	}
	cacheLock.Unlock()

	entry.once.Do(func() {
		raw, err := os.ReadFile(path)
		if err != nil {
			entry.err = fmt.Errorf("failed to read %s: %w", path, err)
			return
		}
		var parsed []T
		if err := json.Unmarshal(raw, &parsed); err != nil {
			entry.err = fmt.Errorf("failed to parse %s: %w", path, err)
			return
		}
		entry.data = parsed
	})

	if entry.err != nil {
		return nil, entry.err
	}

	return entry.data.([]T), nil
}
