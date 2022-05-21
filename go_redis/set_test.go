package go_redis

import "testing"

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
