package go_redis

import "testing"

func TestPFAdd(t *testing.T) {
	key := "test:hll"
	res := client.PFAdd(ctx, key, "a", "b", "c", "d", "e", "f", "g")
	t.Log(res)
	t.Log(client.PFCount(ctx, key))
}

func TestPFCount(t *testing.T) {
	key1 := "hll"
	key2 := "some-other-hl"
	client.PFAdd(ctx, key1, "foo", "bar", "zap")
	client.PFAdd(ctx, key1, "zap", "zap", "zap")
	client.PFAdd(ctx, key1, "foo", "bar")
	t.Log(client.PFCount(ctx, key1))
	client.PFAdd(ctx, key2, 1, 2, 3)
	t.Log(client.PFCount(ctx, key1, key2))
}

func TestPFMerge(t *testing.T) {
	key1 := "hll"
	key2 := "hll2"
	key3 := "hll3"
	client.PFAdd(ctx, key1, "foo", "bar", "zap", "a")
	client.PFAdd(ctx, key2, "a", "b", "c", "foo")
	client.PFMerge(ctx, key3, key1, key2)
	t.Log(client.PFCount(ctx, key3))
}
