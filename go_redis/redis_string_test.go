package go_redis

import (
	"fmt"
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

func TestIncr(t *testing.T) {
	client.Set(ctx, "counter", "1", 0)
	client.Incr(ctx, "counter")
	t.Log(client.Get(ctx, "counter"))
}

func TestIncrBy(t *testing.T) {
	client.IncrBy(ctx, "counter", 10)
	t.Log(client.Get(ctx, "counter"))
}

func TestIncrByFloat(t *testing.T) {
	const key = "test:incrByFloat:key"
	client.Set(ctx, key, 10.5, 0)
	res := client.IncrByFloat(ctx, key, 0.1)
	t.Log(res)
	res = client.IncrByFloat(ctx, key, -5)
	t.Log(res)
	client.Set(ctx, key, 5.0e3, 0)
	res = client.IncrByFloat(ctx, key, 2.0e2)
	t.Log(res)
}

// redis 7.0 新增 api，求字符串的最长相同子序列，当时和最长相同子序列算法要求不同，redis中不要求对比的字符串中的字符是连续的
// Time complexity: O(N*M) where N and M are the lengths of s1 and s2
// 优势是可以很方便的查看两个字符串的相似程度
// 例如使用字符串表示DNA序列，LCS 可以很方便的表示出两条DNA序列的相似程度
// 如果表示用户编辑的文本，LCS 可以方便比较出新旧文档的差异
// 具体可以查看文档 具体可以查看文档 https://redis.io/commands/lcs/
func TestLcs(t *testing.T) {
	client.MSet(ctx, "key1", "ohmytext", "key2", "mynewtext")
	// go-redis v8.11.5 客户端暂时不支持该命令
	//res := client.Lcs(ctx, "key1", "key2")
	//t.Log(res) // "mytext"
}

func TestMGet(t *testing.T) {
	const key1 = "test:mGet:key1"
	const key2 = "test:mGet:key2"
	client.Set(ctx, key1, "Hello", 0)
	client.Set(ctx, key2, "World", 0)
	res := client.MGet(ctx, key1, key2, "noExistString")
	t.Log(res)
}

func TestMSet(t *testing.T) {
	const key1 = "test:mSet:key1"
	const key2 = "test:mSet:key2"
	client.MSet(ctx, key1, "Hello", key2, "World")
	res := client.Get(ctx, key1)
	t.Log(res)
	res = client.Get(ctx, key2)
	t.Log(res)
}

// 当且仅当批量设置的key的值都不存在，才可以批量设置成功
//  MSETNX 返回值是1,所有的值设置成功
//  MSETNX 返回值是0，设置失败（至少有一个值之前已经存在了）
// 在 go-redis 客户端中 成功返回的是 true 失败返回的是 false
func TestMSetNx(t *testing.T) {
	defer func() {
		client.Del(ctx, "key1")
		client.Del(ctx, "key2")
		client.Del(ctx, "key3")
	}()
	res := client.MSetNX(ctx, "key1", "Hello", "key2", "there")
	t.Log(res)
	res = client.MSetNX(ctx, "key2", "new", "key3", "World")
	t.Log(res)
	res2 := client.MGet(ctx, "key1", "key2", "key3")
	t.Log(res2)
}

func TestSetEX(t *testing.T) {
	const key = "test:setEx:key"
	client.SetEX(ctx, key, "Hello", time.Second*10)
	t.Log(client.TTL(ctx, key))
	t.Log(client.Get(ctx, key))
}

// https://redis.io/commands/set/ 在 redis 官网中set命令的介绍中有如下介绍
// The command SET resource-name anystring NX EX max-lock-time is a simple way to implement a locking system with Redis.
// 可以使用SetNx 很方便的实现 Redis 锁
func TestSetNX(t *testing.T) {
	res := client.SetNX(ctx, "test:str:nx", "new_value1", 0)
	t.Log(res)
	res2 := client.SetNX(ctx, "test:str:nx", "new_value2", 0)
	t.Log(res2)
	t.Log(client.Get(ctx, "test:str:nx"))
}

func TestGetLock(t *testing.T) {
	go GetLock("zhang san")
	go GetLock("li si")
	go GetLock("xiao ming")
	time.Sleep(time.Second)
	client.Del(ctx, "mutex")
}

func GetLock(user string) {
	lock := "mutex"
	res := client.SetNX(ctx, lock, user, 0)
	if res.Val() == false {
		owner := client.Get(ctx, lock)
		fmt.Printf("user %s get lock fail, lock owner is %s \n", user, owner.Val())
	} else {
		fmt.Printf("i get a lock, i am %s \n", user)
	}
}

var getLockWg sync.WaitGroup

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

//随机设置某些位置的字符串
//第一次设置时候需要开辟内存空间，第二次不需要了（如果在旧的offset上修改，没有增加新的就不需要）
// 2010 MacBook Pro 上测试 512MB 大概需要300ms 128MB大概需要80ms 32MB 大概需要30ms 8MB 大概需要8ms
func TestSetRange(t *testing.T) {
	const (
		key  = "test:setRange:key1"
		key2 = "test:setRange:key2"
	)
	client.Set(ctx, key, "Hello World", 0)
	client.SetRange(ctx, key, 6, "Redis")
	t.Log(client.Get(ctx, key))
	client.SetRange(ctx, key2, 6, "Redis")
	t.Log(client.Get(ctx, key2))
}

func TestStrLen(t *testing.T) {
	key := "test:strLen:key"
	client.Set(ctx, key, "Hello World", 0)
	t.Log(client.StrLen(ctx, key))
	t.Log(client.StrLen(ctx, "noExitString"))
}
