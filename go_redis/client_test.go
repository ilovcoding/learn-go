package go_redis

import (
	"testing"
	"time"
)

func TestExampleClientSync(t *testing.T) {
	rdb1 := RedisClient()
	rdb2 := RedisClient()
	t.Log(rdb1 == rdb2)
}

func TestClientSave(t *testing.T) {
	rdb := RedisClient()
	key := "local:test"
	rdb.Set(ctx, key, time.Now().String(), 0)
	t.Log(rdb.Get(ctx, key))
	rdb.Del(ctx, key)
}
