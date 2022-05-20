package go_redis

import (
	"testing"
)

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

func TestHGet(t *testing.T) {
	client.HSet(ctx, key, "field1", "foo")
	res := client.HGet(ctx, key, "field1")
	t.Log(res)
	res = client.HGet(ctx, key, "field2")
	t.Log(res)
}

func TestHGetAll(t *testing.T) {
	client.HSet(ctx, key, "field1", "Hello")
	client.HSet(ctx, key, "field2", "World")
	res := client.HGetAll(ctx, key)
	t.Log(res)
}

func TestHIncrBy(t *testing.T) {
	client.HSet(ctx, key, "field", 5)
	res := client.HIncrBy(ctx, key, "field", 1)
	t.Log(res)
}

func TestHIncrByFloat(t *testing.T) {
	client.HSet(ctx, key, "field", 10.50)
	res := client.HIncrByFloat(ctx, key, "field", 0.1)
	t.Log(res)
}

func TestHKeys(t *testing.T) {
	defer func() {
		client.Del(ctx, key)
	}()
	client.HSet(ctx, key, "field1", "Hello")
	client.HSet(ctx, key, "field2", "World")
	t.Log(client.HKeys(ctx, key))
}

func TestHLen(t *testing.T) {
	client.HSet(ctx, key, "field1", "Hello")
	client.HSet(ctx, key, "field2", "World")
	res := client.HLen(ctx, key)
	t.Log(res)
}

func TestHMGet(t *testing.T) {
	res := client.HMGet(ctx, key, "field1", "field2", "noField")
	t.Log(res)
}

// HRANDFIELD key [ count [WITHVALUES]]
// 从 6.2.0 版本才开始使用 如果 count 为正整数则随机返回一些不重复的field，如果count为负数，随机返回一些可以重复的field
// 返回的个数为count的绝对值。
func TestHRandField(t *testing.T) {
	defer func() {
		client.Del(ctx, key)
	}()
	client.HMSet(ctx, key, "heads", "obverse", "tails", "reverse", "edge", nil)
	res := client.HRandField(ctx, key, 1, false)
	t.Log(res)
	res = client.HRandField(ctx, key, 1, true)
	t.Log(res)
	res = client.HRandField(ctx, key, -5, false)
	t.Log(res)
	res = client.HRandField(ctx, key, -5, true)
	t.Log(res)
}

func TestHSet(t *testing.T) {
	client.HSet(ctx, key, "field1", "Hello")
	res := client.HGet(ctx, key, "field1")
	t.Log(res)
}

func TestHSetNX(t *testing.T) {
	defer func() {
		client.Del(ctx, key)
	}()
	client.HSetNX(ctx, key, "field", "Hello")
	client.HSetNX(ctx, key, "field", "World")
	res := client.HGet(ctx, key, "field")
	t.Log(res)
}

func TestHValS(t *testing.T) {
	client.HSet(ctx, key, "field1", "Hello")
	client.HSet(ctx, key, "field2", "World")
	res := client.HVals(ctx, key)
	t.Log(res)
	res = client.HVals(ctx, "noKwy")
	t.Log(res)
}
