package go_redis

import "testing"

const setKey = "test:set:key"

func TestSAdd(t *testing.T) {
	client.SAdd(ctx, setKey, "Hello")
	client.SAdd(ctx, setKey, "World")
	client.SAdd(ctx, setKey, "World")
	res := client.SMembers(ctx, setKey)
	t.Log(res)
}
