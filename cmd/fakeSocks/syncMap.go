package fakeSocks

import "sync"

type counters struct {
	mx sync.RWMutex
	m  map[string]uint8 //255
}

func newCounters() *counters {
	return &counters{
		m: make(map[string]uint8),
	}
}

func (c *counters) Load(key string) int {
	c.mx.RLock()
	defer c.mx.RUnlock()
	val, _ := c.m[key]
	return int(val)
}

func (c *counters) Inc(key string) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m[key]++
}
