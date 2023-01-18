package storage

import (
	"encoding/json"
	"os"
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

	val := c.m[key]
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
func (c *storageMap) WriteToFile(fileMap string) error {
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

func (c *storageMap) Increment(key string) {
}
