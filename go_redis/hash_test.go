package go_redis

import "testing"

const key = "test:hash:key"

func TestHDel(t *testing.T) {
	client.HSet(ctx, key, "field1", "foo")
	res := client.HDel(ctx, key, "field1")
	t.Log(res)
	res = client.HDel(ctx, key, "field2")
	t.Log(res)
}

func TestHExists(t *testing.T) {
	client.HSet(ctx, key, "field1", "foo")
	res := client.HExists(ctx, key, "field1")
	t.Log(res) // truw
	res = client.HExists(ctx, key, "field2")
	t.Log(res) // false
}
