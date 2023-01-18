package storage

import (
	"sync"
)

type counters struct {
	mx sync.RWMutex
	m  map[string]uint
}

func NewCountersMap() *counters {
	return &counters{
		m: make(map[string]uint),
	}
}

func (c *counters) Load(key string) int64 {
	c.mx.RLock()
	defer c.mx.RUnlock()

	val := c.m[key]
	return int64(val)
}

func (c *counters) Increment(key string) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m[key]++
}

func (c *counters) Save(key string, v int64) {
}
func (c *counters) ReadFromFile(fileMap string) error {
	return nil
}
func (c *counters) WriteToFile(fileMap string) error {
	return nil
}
