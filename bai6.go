package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func Bai6(connect *redis.Client, ctx context.Context) {
	err := connect.Set(ctx, "demo_key", time.Now().Unix(), 10*time.Second).Err()
	if err != nil {
		fmt.Println(err)
	}

	val, err := connect.Get(ctx, "demo_key").Result()
	if err == redis.Nil {
		fmt.Println("Key doesn't exist!")
	} else {
		fmt.Println("Value: ", val)
	}
	time.Sleep(12 * time.Second)

	val1, err := connect.Get(ctx, "demo_key").Result()

	if err == redis.Nil {
		err = connect.Set(ctx, "demo_key", time.Now().Unix(), 0).Err()
		if err != nil {
			fmt.Println(err)
		}
		val1, err = connect.Get(ctx, "demo_key").Result()
		if err == redis.Nil {
			fmt.Println("Key doesn't exist!")
		} else {
			fmt.Println("Value after 12s: ", val1)
		}
	} else {
		fmt.Println("Value: ", val1)
	}
}
