// nolint
package main

import (
	"sync"
)

type counters struct {
	mx sync.RWMutex
	m  map[string]int
}

var loginAttemptsIP = &counters{
	m: make(map[string]int),
}

func (c *counters) Load(key string) int {
	c.mx.RLock()
	defer c.mx.RUnlock()

	val, _ := c.m[key]
	return val
}

func (c *counters) Increment(key string) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m[key]++
}

func (c *counters) Unblock(key string) {
	c.mx.RLock()
	biggest := c.m[key] > 0
	c.mx.RUnlock()

	if biggest {
		c.mx.Lock()
		c.m[key] = -30 << 30
		c.mx.Unlock()
	}
}
