package storage

import (
	"encoding/json"
	"os"
	"sync"
)

// SyncMap provides a thread-safe map with file persistence capabilities
type SyncMap struct {
	mx sync.RWMutex
	m  map[string]int64
}

// NewSyncMap creates a new thread-safe map instance
func NewSyncMap() *SyncMap {
	return &SyncMap{
		m: make(map[string]int64),
	}
}

// Increment increases the value for the given key by 1
func (c *SyncMap) Increment(key string) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m[key]++
}

// Load retrieves the value for the given key
func (c *SyncMap) Load(key string) int64 {
	c.mx.RLock()
	defer c.mx.RUnlock()

	val := c.m[key]
	return val
}

// Save sets the value for the given key
func (c *SyncMap) Save(key string, v int64) {
	c.mx.Lock()
	defer c.mx.Unlock()

	c.m[key] = v
}

// ReadFromFile loads the map data from a JSON file
func (c *SyncMap) ReadFromFile(fileMap string) error {
	// Check if file exists
	if _, err := os.Stat(fileMap); os.IsNotExist(err) {
		return err
	}

	f, err := os.Open(fileMap)
	if err != nil {
		return err
	}
	defer f.Close()

	fileInfo, err := os.Stat(fileMap)
	if err != nil {
		os.Remove(fileMap)
		return err
	}

	buf := make([]byte, fileInfo.Size())
	_, err = f.Read(buf)
	if err != nil {
		os.Remove(fileMap)
		return err
	}

	// Use write lock since we're modifying the map
	c.mx.Lock()
	defer c.mx.Unlock()

	err = json.Unmarshal(buf, &c.m)
	if err != nil {
		os.Remove(fileMap)
		return err
	}

	return nil
}

// WriteToFile saves the map data to a JSON file
func (c *SyncMap) WriteToFile(fileMap string) error {
	c.mx.Lock()
	defer c.mx.Unlock()

	buf, err := json.Marshal(c.m)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(fileMap, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(buf)
	if err != nil {
		return err
	}

	if err = f.Sync(); err != nil {
		return err
	}

	return nil
}
