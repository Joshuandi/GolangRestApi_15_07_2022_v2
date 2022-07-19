package cache

import (
	"GolangRestApi_15_07_2022_v2/model"
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
)

type RedisCacheInterface interface {
	Set(key string, users *model.Users)
	Get(key string) *model.Users
}

type RedisCache struct {
	Host string
	Dbr  int
	Exp  time.Duration
}

func NewRedisCache(host string, dbr int, exp time.Duration) RedisCacheInterface {
	return &RedisCache{
		Host: host,
		Dbr:  dbr,
		Exp:  exp,
	}
}
func (c *RedisCache) ClientCon() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Host,
		Password: "",
		DB:       c.Dbr,
	})
}

func (c *RedisCache) Set(key string, users *model.Users) {
	cl := c.ClientCon()
	json, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}
	cl.Set(key, json, c.Exp*time.Second)
}

func (c *RedisCache) Get(key string) *model.Users {
	cl := c.ClientCon()
	val, err := cl.Get(key).Result()
	if err != nil {
		panic(err)
	}

	value := model.Users{}
	err = json.Unmarshal([]byte(val), &value)
	if err != nil {
		panic(err)
	}

	return &value
}
