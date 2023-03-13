package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})

	err := rdb.HSet(ctx, "hash", "field1", "value1").Err()
	if err != nil {
		fmt.Println("redis set hash error=", err)
	}

	err = rdb.HSet(ctx, "hash", "field2", "value2").Err()
	if err != nil {
		fmt.Println("redis set hash error=", err)
	}

	str, err := rdb.HGet(ctx, "hash", "field1").Result()
	if err != nil {
		fmt.Println("get hash field1 err=", err)
	}
	fmt.Println("hash field1 value=", str)

	vals, err := rdb.HGetAll(ctx, "hash").Result()
	if err != nil {
		fmt.Println("get hash all fields error=", err)
	}
	fmt.Println("hash all values=", vals)
}
