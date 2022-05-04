package go_redis

import (
	"sync"
	"testing"
	"time"
)

// 二进制安全
func TestBinarySecurity(t *testing.T) {
	client.Set(ctx, "", "空字符串", 0)
	t.Log(client.Get(ctx, ""))
	client.Set(ctx, "\\0", "zero", 0)
	t.Log(client.Get(ctx, "\\0"))
}

func TestAppend(t *testing.T) {
	const key = "test:append:key"
	defer client.Del(ctx, key)
	res := client.Append(ctx, key, "Hello")
	t.Log(res)
	res = client.Append(ctx, key, " World")
	t.Log(res.Result())
	t.Log(client.Get(ctx, key))
	// Append 裁剪字符串的功能
	res2 := client.GetRange(ctx, key, 3, 4)
	t.Log(res2)
}

func TestDecr(t *testing.T) {
	const key = "test:decr:key"
	defer client.Del(ctx, key)
	res := client.Decr(ctx, key)
	t.Log(res)
	client.Set(ctx, key, 10, 0)
	res = client.Decr(ctx, key)
	t.Log(res)
	client.Set(ctx, key, "234293482390480948029348230948", 0)
	res = client.Decr(ctx, key)
	t.Log(res)
}

func TestGet(t *testing.T) {
	const key = "test:get:key"
	res := client.Get(ctx, key)
	t.Log(res)
	res2 := client.Set(ctx, key, "Hello", 0)
	t.Log(res2)
	res = client.Get(ctx, key)
	t.Log(res)
}

func TestDecrBy(t *testing.T) {
	const key = "test:decr:key"
	res := client.DecrBy(ctx, key, 3)
	t.Log(res)
}

func TestGetDel(t *testing.T) {
	const key = "test:getDel:key"
	client.Set(ctx, key, 10, 0)
	res := client.GetDel(ctx, key)
	t.Log(res)
	client.HSet(ctx, key, "key1", "value1")
	//	 WRONGTYPE Operation against a key holding the wrong kind of value
	// 当且仅当 value 的 type 是 string 的时候才能使用 GetDel 否则会返回错误
	res2 := client.GetDel(ctx, key)
	t.Log(res2)
}

func TestGetEx(t *testing.T) {
	const key = "test:getEx:key"
	client.Set(ctx, key, "Hello", 0)
	// -1纳秒 永不过期
	t.Log(client.TTL(ctx, key).Val())
	res := client.GetEx(ctx, key, time.Second*60)
	t.Log(res)
	// 1分0秒 60s 过期
	t.Log(client.TTL(ctx, key).Val())
}

func TestGetRange(t *testing.T) {
	const key = "test:getRange:key"
	client.Set(ctx, key, "This is a string", 0)
	res := client.GetRange(ctx, key, 0, 3)
	t.Log(res)
	res = client.GetRange(ctx, key, -3, -1)
	t.Log(res)
	res = client.GetRange(ctx, key, 0, -1)
	t.Log(res)
	res = client.GetRange(ctx, key, 10, 100)
	t.Log(res)
}

func TestNX(t *testing.T) {
	res := client.SetNX(ctx, "test:str:nx", "new_value1", 0)
	t.Log(res)
	res2 := client.SetNX(ctx, "test:str:nx", "new_value2", 0)
	t.Log(res2)
	t.Log(client.Get(ctx, "test:str:nx"))
}

var getLockWg sync.WaitGroup

func TestGetLock(t *testing.T) {
	go GetLock("zhang san")
	go GetLock("li si")
	go GetLock("xiao ming")
	time.Sleep(time.Second)
	client.Del(ctx, "mutex")
}

func TestIncr(t *testing.T) {
	client.Set(ctx, "counter", "1", 0)
	client.Incr(ctx, "counter")
	t.Log(client.Get(ctx, "counter"))
	client.IncrBy(ctx, "counter", 10)
	t.Log(client.Get(ctx, "counter"))
}

func TestGetSet(t *testing.T) {
	res := client.GetSet(ctx, "counter", -1)
	t.Log(res.Val())
	t.Log(client.Get(ctx, "counter"))
}

func TestGetLockV2(t *testing.T) {
	getLockWg.Add(1)
	go func() {
		GetLock("zhang san")
		getLockWg.Done()
	}()
	getLockWg.Add(1)
	go func() {
		GetLock("li si")
		getLockWg.Done()
	}()
	getLockWg.Add(1)
	go func() {
		GetLock("xiao ming")
		getLockWg.Done()
	}()
	getLockWg.Wait()
	client.Del(ctx, "mutex")
}
