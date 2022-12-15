package syncMap

import "sync"

type storageMap struct {
	mx sync.RWMutex
	m  map[string]int64
}

func NewStorageMap() *storageMap {
	return &storageMap{
		m: make(map[string]int64),
	}
}

func (c *storageMap) Load(key string) int64 {
	c.mx.RLock()
	defer c.mx.RUnlock()
	val, _ := c.m[key]
	return val
}

func (c *storageMap) Save(key string, v int64) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	c.m[key] = v
}

func (c *storageMap) Increment(key string) {
}
