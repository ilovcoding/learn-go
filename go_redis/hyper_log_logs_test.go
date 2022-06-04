package go_redis

import "testing"

func TestPFAdd(t *testing.T) {
	key := "test:hll"
	res := client.PFAdd(ctx, key, "a", "b", "c", "d", "e", "f", "g")
	t.Log(res)
	t.Log(client.PFCount(ctx, key))
}
