package main

import (
	"github.com/go-redis/redis/v8"
)

var connRedis = Redis()

func Redis() *redis.Client {
	redis := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return redis
}
