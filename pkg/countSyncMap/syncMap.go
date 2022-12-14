package countSyncMap

import "sync"

type Counters struct {
	mx sync.RWMutex
	m  map[string]uint8 //255
}

func NewCounters() *Counters {
	return &Counters{
		m: make(map[string]uint8),
	}
}

func (c *Counters) Load(key string) int {
	c.mx.RLock()
	defer c.mx.RUnlock()
	val, _ := c.m[key]
	return int(val)
}

func (c *Counters) Inc(key string) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m[key]++
}
