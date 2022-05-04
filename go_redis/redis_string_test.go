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
	client.Append(ctx, key, "Hello")
	client.Append(ctx, key, " World")
	t.Log(client.Get(ctx, key))
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
