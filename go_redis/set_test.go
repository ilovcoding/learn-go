package go_redis

import (
	"github.com/go-redis/redis/v8"
	"testing"
)

const setKey = "test:set:key"

func TestSAdd(t *testing.T) {
	client.SAdd(ctx, setKey, "Hello")
	client.SAdd(ctx, setKey, "World")
	client.SAdd(ctx, setKey, "World")
	res := client.SMembers(ctx, setKey)
	t.Log(res)
}

// 返回 Set 中含有元素个数
func TestSCard(t *testing.T) {
	res := client.SCard(ctx, setKey)
	t.Log(res)
}

func TestSDiff(t *testing.T) {
	defer func() {
		client.Del(ctx, "key1", "key2", "key3")
	}()
	client.SAdd(ctx, "key1", "a", "b", "c", "d")
	client.SAdd(ctx, "key2", "c")
	client.SAdd(ctx, "key3", "a", "c", "e")
	res := client.SDiff(ctx, "key1", "key2", "key3")
	t.Log(res)
}

// SDIFFSTORE destination key [key ...]
func TestSDiffStore(t *testing.T) {
	defer func() {
		client.Del(ctx, "key1", "key2", "key")
	}()
	client.SAdd(ctx, "key1", "a", "b", "c")
	client.SAdd(ctx, "key2", "c", "d", "e")
	client.SDiffStore(ctx, "key", "key1", "key2")
	res := client.SMembers(ctx, "key")
	t.Log(res)
}

// 求集合交集 O(M*N)
func TestSInter(t *testing.T) {
	defer func() {
		client.Del(ctx, "key1", "key2", "key3")
	}()
	client.SAdd(ctx, "key1", "a", "b", "c", "d")
	client.SAdd(ctx, "key2", "c")
	client.SAdd(ctx, "key3", "a", "c", "e")
	res := client.SInter(ctx, "key1", "key2", "key3")
	t.Log(res)
}

func TestSInterStore(t *testing.T) {
	client.SAdd(ctx, "key1", "a", "b", "c")
	client.SAdd(ctx, "key2", "c", "d", "e")
	res := client.SInterStore(ctx, key, "key1", "key2")
	t.Log(res)
	res2 := client.SMembers(ctx, key)
	t.Log(res2)
}

// O(1)
func TestSIsMember(t *testing.T) {
	client.SAdd(ctx, key, "one")
	res := client.SIsMember(ctx, key, "one")
	t.Log(res)
	res = client.SIsMember(ctx, key, "two")
	t.Log(res)
}

func TestSMIsMember(t *testing.T) {
	client.SAdd(ctx, setKey, "one")
	res := client.SMIsMember(ctx, setKey, "one", "notAMember")
	t.Log(res)
}

func TestSMove(t *testing.T) {
	defer func() {
		client.Del(ctx, setKey)
	}()
	client.SAdd(ctx, setKey, "one")
	client.SAdd(ctx, setKey, "two")
	client.SAdd(ctx, "myOtherSet", "three")
	client.SMove(ctx, setKey, "myOtherSet", "two")
	t.Log(client.SMembers(ctx, setKey))
	t.Log(client.SMembers(ctx, "myOtherSet"))

}

func TestSPop(t *testing.T) {
	key := "test:key:sPop"
	client.SAdd(ctx, key, "one", "two", "three")
	res := client.SPop(ctx, key)
	t.Log(res)
	res2 := client.SMembers(ctx, key)
	t.Log(res2)
	client.SAdd(ctx, key, "four", "five")
	res3 := client.SPopN(ctx, key, 3)
	t.Log(res3)
	res2 = client.SMembers(ctx, key)
	t.Log(res2)
}

func TestSRandMember(t *testing.T) {
	key := "test:sRandMember:key"
	client.SAdd(ctx, key, "one", "two", "three")
	res := client.SRandMember(ctx, key)
	t.Log(res)
	res2 := client.SRandMemberN(ctx, key, 2)
	t.Log(res2)
	res3 := client.SRandMemberN(ctx, key, -5)
	t.Log(res3)
}

func TestSRem(t *testing.T) {
	key := "test:SRem:key"
	client.SAdd(ctx, key, "one", "two", "three")
	res := client.SRem(ctx, key, "one")
	t.Log(res)
	res2 := client.SRem(ctx, key, "four")
	t.Log(res2)
	res3 := client.SMembers(ctx, key)
	t.Log(res3)
}

func TestSUnion(t *testing.T) {
	key1 := "test:SUnion:key1"
	key2 := "test:SUnion:key2"
	client.SAdd(ctx, key1, "one", "two", "three")
	client.SAdd(ctx, key2, "one", "four", "five")
	res := client.SUnion(ctx, key1, key2)
	t.Log(res)
}

func TestSUnionStore(t *testing.T) {
	key1 := "test:SUnionStore:key1"
	key2 := "test:SUnionStore:key2"
	key := "test:SUnionStore:key"
	client.SAdd(ctx, key1, "one", "two", "three")
	client.SAdd(ctx, key2, "one", "four", "five")
	client.SUnionStore(ctx, key, key1, key2)
	res := client.SMembers(ctx, key)
	t.Log(res)
}

// iterate 模式，遍历redis中的key可用于翻页等
func TestScan(t *testing.T) {
	res := client.Scan(ctx, 0, "", 0)
	t.Log(res)
	page, cursor, err := res.Result()
	t.Log(page)
	t.Log(cursor)
	t.Log(err)
	res = client.Scan(ctx, 28, "", 0)
	t.Log(res)
	page, cursor, err = res.Result()
	t.Log(page)
	t.Log(cursor)
	t.Log(err)
	res = client.Scan(ctx, 38, "", 0)
	t.Log(res)
	page, cursor, err = res.Result()
	t.Log(page)
	t.Log(cursor)
	t.Log(err)
	res = client.Scan(ctx, 17, "", 0)
	t.Log(res)
	page, cursor, err = res.Result()
	t.Log(page)
	t.Log(cursor)
	t.Log(err)
	res = client.Scan(ctx, 29, "", 0)
	t.Log(res)
	page, cursor, err = res.Result()
	t.Log(page)
	t.Log(cursor)
	t.Log(err)
}

// iterate 模式迭代 某一个 set 中 含有的元素。
func TestSScan(t *testing.T) {
	client.SAdd(ctx, setKey, "foo", "feelsgood", "foobar", "0", "1", "2", 3, 4, 5, "a", "b")
	res := client.SScan(ctx, setKey, 0, "f*", 11)
	t.Log(res)
	page, cursor, err := res.Result()
	t.Log(page)
	t.Log(cursor)
	t.Log(err)
	res = client.SScan(ctx, setKey, 0, "f*", 11)
	t.Log(res)
	page, cursor, err = res.Result()
	t.Log(page)
	t.Log(cursor)
	t.Log(err)
}

func TestTypeScan(t *testing.T) {
	client.ZAdd(ctx, "zKey1", &redis.Z{Score: 100, Member: "value1"})
	client.ZAdd(ctx, "zKey2", &redis.Z{Score: 0, Member: "value2"})
	t.Log(client.Type(ctx, "zKey1"))
	res := client.ScanType(ctx, 0, "zK*", 0, "zset")
	t.Log(res.Result())
	res = client.ScanType(ctx, 28, "zK*", 0, "zset")
	t.Log(res.Result())
	res = client.ScanType(ctx, 38, "zK*", 0, "zset")
	t.Log(res.Result())
	res = client.ScanType(ctx, 17, "zK*", 0, "zset")
	t.Log(res.Result())
	res = client.ScanType(ctx, 21, "zK*", 0, "zset")
	t.Log(res.Result())
	res = client.ScanType(ctx, 47, "zK*", 0, "zset")
	t.Log(res.Result())
}
