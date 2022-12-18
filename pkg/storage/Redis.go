package storage

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"io/ioutil"
	"log"
	"strconv"
)

type redisClient struct {
	cl *redis.Client
}

func NewRedis() *redisClient {
	return &redisClient{
		cl: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "redis",
			DB:       0,
		}),
	}
}

func (c *redisClient) Load(key string) int64 {
	val, err := c.cl.Get(key).Result()
	if err != nil {
		log.Println(err)
		return 0
	}
	end, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		log.Println(err)
	}
	return end
}

func (c *redisClient) Save(key string, val int64) {

	err := c.cl.Set(key, val, 0).Err()

	if err != nil {
		log.Println(err)
	}

}

func (c *redisClient) Increment(key string) {
	err := c.cl.Incr(key).Err()
	if err != nil {
		log.Println(err)
	}
}

func (c *redisClient) ReadFromFile(fileMap string) error {
	m := new(map[string]int64)
	buf, err := ioutil.ReadFile(fileMap)
	if err == nil {
		err = json.Unmarshal(buf, m)
	}
	if err != nil {
		return err
	}
	err = c.cl.FlushAll().Err()
	if err != nil {
		return err
	}

	for k, v := range *m {

		err = c.cl.Set(k, v, 0).Err()
		if err != nil {
			return err
		}
	}

	return err
}

func (c *redisClient) WriteToFile(fileMap string) error {
	m := make(map[string]int64)

	arr, err := c.cl.Keys("*").Result()
	if err != nil {
		log.Println(err)
		return err
	}

	for _, st := range arr {
		m[st] = c.Load(st)
	}

	buf, err := json.Marshal(m)
	if err == nil {

		err = ioutil.WriteFile(fileMap, buf, 0644)
	}
	return err
}

func (c *redisClient) Close() {
	c.cl.Close()
}
