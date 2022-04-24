package go_redis

import (
	"sync"
	"testing"
	"time"
)

func TestBinarySecurity(t *testing.T) {
	client.Set(ctx, "", "空字符串", 0)
	t.Log(client.Get(ctx, ""))
	client.Set(ctx, "\\0", "zero", 0)
	t.Log(client.Get(ctx, "\\0"))
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
