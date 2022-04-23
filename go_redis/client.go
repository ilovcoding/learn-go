package go_redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var client *redis.Client
var rdb = client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "192.168.0.113:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("redis init error", err)
	}
}

func RedisClient() *redis.Client {
	return client
}
