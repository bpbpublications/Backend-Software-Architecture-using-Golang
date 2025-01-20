package main

import (
    "context"
    "github.com/go-redis/redis/v8"
    "log"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })

    // Producer
    _, err := rdb.XAdd(ctx, &redis.XAddArgs{
        Stream: "mystream",
        Values: map[string]interface{}{"message": "hello"},
    }).Result()
    if err != nil {
        log.Fatal(err)
    }

    // Consumer
    result, err := rdb.XRead(ctx, &redis.XReadArgs{
        Streams: []string{"mystream", "0"},
        Count:   1,
        Block:   0,
    }).Result()
    if err != nil {
        log.Fatal(err)
    }
    for _, stream := range result {
        for _, message := range stream.Messages {
            log.Printf("Received message: %v", message.Values)
        }
    }
}
