package go_redis

import (
	"github.com/go-redis/redis/v8"
	"testing"
)

func TestExampleClient(t *testing.T) {
	var rdb1 *redis.Client
	var rdb2 *redis.Client
	go func() {
		rdb1 = ExampleClient()
	}()
	go func() {
		rdb2 = ExampleClient()
	}()
	t.Log(rdb1 == rdb2)
}

func TestExampleClientAsync(t *testing.T) {
	for {
		TestExampleClient(t)
	}
}

func TestExampleClientSync(t *testing.T) {
	rdb1 := ExampleClient()
	rdb2 := ExampleClient()
	t.Log(rdb1 == rdb2)
}
