package go_redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client

func ExampleClient() *redis.Client {
	if rdb != nil {
		return rdb
	}
	rdb = redis.NewClient(&redis.Options{
		Addr:     "192.168.0.113:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}
