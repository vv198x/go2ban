package storage

import (
	"encoding/json"
	"io/ioutil"
	"sync"
)

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
	c.mx.Lock()
	defer c.mx.Unlock()

	c.m[key] = v
}

func (c *storageMap) ReadFromFile(fileMap string) error {
	c.mx.RLock()
	defer c.mx.RUnlock()

	buf, err := ioutil.ReadFile(fileMap)
	if err == nil {

		err = json.Unmarshal(buf, &c.m)
	}
	return err
}
func (c *storageMap) WriteToFile(fileMap string) error {
	c.mx.Lock()
	defer c.mx.Unlock()

	buf, err := json.Marshal(c.m)
	if err == nil {

		err = ioutil.WriteFile(fileMap, buf, 0644)
	}
	return err
}

func (c *storageMap) Increment(key string) {
}
