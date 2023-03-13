package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})

	//set
	err := rdb.Set(ctx, "key", "value123", 0).Err()
	if err != nil {
		fmt.Println("redis set key-value err=", err)
	}

	//get
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		fmt.Println("redis get value err=", err)
	}
	fmt.Println("value=", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		fmt.Println("redis get value err=", err)
	} else {
		fmt.Println("value2=", val2)
	}
}
