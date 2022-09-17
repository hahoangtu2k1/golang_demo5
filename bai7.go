package main

import (
	"bufio"
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

func Bai7(connect *redis.Client, ctx context.Context) {
	file, err := os.Open("name.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		connect.LPush(ctx, "userName", scanner.Text())
	}
}
