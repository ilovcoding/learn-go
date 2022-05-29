package go_redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var client *redis.Client
var rdb = client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "192.168.140.134:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic("redis init error")
	}
}

func RedisClient() *redis.Client {
	return client
}
