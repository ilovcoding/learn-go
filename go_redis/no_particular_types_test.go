/**
没有指明特定类型的操作指令，一般可以操作任意 Key
*/

package go_redis

import (
	"testing"
	"time"
)

// 判断 key 是否存在, 不存在返回0 存在返回 1
func TestExists(t *testing.T) {
	key := "test:no_particular_types:key"
	res := client.Exists(ctx, key)
	t.Log(res) // 0
	client.Set(ctx, key, "value", 0)
	res = client.Exists(ctx, key)
	t.Log(res) // 1
}

// 删除 key，数据库里原本有key返回1，数据库里本来就没有key返回0
func TestDel(t *testing.T) {
	key := "test:no_particular_types:key"
	res := client.Del(ctx, key)
	t.Log(res) // 1
	res = client.Exists(ctx, key)
	t.Log(res) // 0
	res = client.Del(ctx, key)
	t.Log(res) // 0
}

// 返回key中存储对应值的类型
func TestType(t *testing.T) {
	key := "test:no_particular_types:key"
	defer func() {
		client.Del(ctx, key)
	}()
	res := client.Type(ctx, key)
	t.Log(res) // none
	client.Set(ctx, key, "1", 0)
	t.Log(client.Type(ctx, key)) // string
	client.Incr(ctx, key)
	t.Log(client.Type(ctx, key)) // string
	client.Del(ctx, key)
	client.LPush(ctx, key, "A")
	t.Log(client.Type(ctx, key)) // list
}

const testExpireKey = "test:expire:key"

/** 设置超时时间而且时间会记录到磁盘中，因为记录的是日期时间所有服务就算重启停止了，在这过程中超时计算依然存在 **/
func TestKeyExpiration(t *testing.T) {
	client.Set(ctx, testExpireKey, "some-value", 0)
	t.Log(client.TTL(ctx, testExpireKey))
	client.Expire(ctx, testExpireKey, time.Second*5)
}

func TestTTL(t *testing.T) {
	res := client.TTL(ctx, testExpireKey)
	t.Log(client.Get(ctx, testExpireKey))
	t.Log(res)
}

func TestKeys(t *testing.T) {
	client.MSet(ctx, "firstname", "Jack", "lastname", "Stuntman", "age", 35)
	res := client.Keys(ctx, "*name*")
	t.Log(res)
	res = client.Keys(ctx, "a??")
	t.Log(res)
	res = client.Keys(ctx, "*")
	t.Log(res)
}
