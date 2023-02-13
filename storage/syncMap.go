package storage

import (
	"encoding/json"
	"os"
	"sync"
)

type syncMap struct {
	mx sync.RWMutex
	m  map[string]int64
}

func NewSyncMap() *syncMap {
	return &syncMap{
		m: make(map[string]int64),
	}
}

func (c *syncMap) Increment(key string) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m[key]++
}

func (c *syncMap) Load(key string) int64 {
	c.mx.RLock()
	defer c.mx.RUnlock()

	val := c.m[key]
	return val
}

func (c *syncMap) Save(key string, v int64) {
	c.mx.Lock()
	defer c.mx.Unlock()

	c.m[key] = v
}

func (c *syncMap) ReadFromFile(fileMap string) error {
	c.mx.RLock()
	defer c.mx.RUnlock()

	f, err := os.Open(fileMap)
	if err != nil {
		return err
	}
	defer f.Close()

	fileInfo, err := os.Stat(fileMap)
	if err != nil {
		return err
	}

	buf := make([]byte, fileInfo.Size())
	_, err = f.Read(buf)
	if err != nil {
		return err
	}

	err = json.Unmarshal(buf, &c.m)

	return err
}

func (c *syncMap) WriteToFile(fileMap string) error {
	c.mx.Lock()
	defer c.mx.Unlock()

	buf, err := json.Marshal(c.m)
	if err == nil {

		f, errF := os.OpenFile(fileMap, os.O_WRONLY|os.O_CREATE, 0644)
		if errF != nil {
			return errF
		}
		defer f.Close()

		_, errF = f.Write(buf)
		if errF != nil {
			return errF
		}

	}
	return err
}
