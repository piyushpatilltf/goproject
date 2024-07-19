package config

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client

func ConnectToRedis() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "127.0.0.1:6379",
        Password: "",             
    })

    pong, err := rdb.Ping(context.Background()).Result()
    if err != nil {
        panic(err)
    }
    fmt.Println(pong, err)

    Rdb = rdb
}
