package go_redis

import (
	"github.com/go-redis/redis/v8"
	"testing"
)

const bitMapKey = "test:bit_map:key"

// 统计出 bit 集合中 含有1的个数（高电频的个数）, Redis 7.0 增加了 BYTE|BIT 的操作，来指定 start end 的索引类型是 byte 还是 bit. 默认是 byte.
func TestBitCount(t *testing.T) {
	client.Set(ctx, bitMapKey, "foobar", 0)
	res := client.BitCount(ctx, bitMapKey, &redis.BitCount{Start: 0, End: -1})
	t.Log(res)
	// 102  f 1100110
	res = client.BitCount(ctx, bitMapKey, &redis.BitCount{Start: 0, End: 0})
	t.Log(res)
	// 111  f 1101111
	res = client.BitCount(ctx, bitMapKey, &redis.BitCount{Start: 1, End: 1})
	t.Log(res)
	res = client.BitCount(ctx, bitMapKey, &redis.BitCount{Start: 5, End: 30})
	t.Log(res)
}

// 再指定的位置进行一组二进制位的操作 可能不常用
func TestBitField(t *testing.T) {
	defer func() {
		client.Del(ctx, bitMapKey)
	}()
	client.Set(ctx, bitMapKey, 0, 0)
	t.Log(client.Get(ctx, bitMapKey))
	res := client.BitCount(ctx, bitMapKey, &redis.BitCount{Start: 0, End: -1})
	t.Log(res)
	//BITFIELD mykey INCRBY i5 100 1 GET u4 0
	// 0000000001 (i5 5 1 = 1) 0111 (u10 0 = 1)
	res2 := client.BitField(ctx, bitMapKey, "INCRBY", "i5", 5, 1, "GET", "u10", 0)
	t.Log(res2)
	// 0111100001( "i5", 0, 15) (U10 0=1+32+64+128+256=481)
	res2 = client.BitField(ctx, bitMapKey, "INCRBY", "i5", 0, 15, "GET", "u10", 0)
	t.Log(res2)
	res = client.BitCount(ctx, bitMapKey, &redis.BitCount{Start: 0, End: -1})
	t.Log(res)
}

// 在多个key之间进行二进制位操作
func TestBitOP(t *testing.T) {
	key1 := "test:bit_map:key1"
	key2 := "test:bit_map:key2"
	key3 := "test:bit_map:key3"

	defer func() {
		client.Del(ctx, key1, key2, key3)
	}()

	// AND 与
	// 00110001 (49)
	client.Set(ctx, key1, 1, 0)
	// 00110010 (50)
	client.Set(ctx, key2, 2, 0)
	// 00110000 (48) 0
	client.BitOpAnd(ctx, key3, key1, key2)
	t.Log(client.Get(ctx, key3))
	// OR 或
	// 00110011 (51) 3
	client.BitOpOr(ctx, key3, key1, key2)
	t.Log(client.Get(ctx, key3))
	// NOT 非
	// 11001110 (206)
	client.BitOpNot(ctx, key3, key1)
	t.Log(client.GetBit(ctx, key3, 0))
	t.Log(client.GetBit(ctx, key3, 1))
	t.Log(client.GetBit(ctx, key3, 2))
	t.Log(client.GetBit(ctx, key3, 3))
	t.Log(client.GetBit(ctx, key3, 4))
	t.Log(client.GetBit(ctx, key3, 5))
	t.Log(client.GetBit(ctx, key3, 6))
	t.Log(client.GetBit(ctx, key3, 7))
	// XOR 异或
	// 00000011 （3）
	client.BitOpXor(ctx, key3, key1, key2)
	t.Log(client.Get(ctx, key3))
	t.Log(client.GetBit(ctx, key3, 0))
	t.Log(client.GetBit(ctx, key3, 1))
	t.Log(client.GetBit(ctx, key3, 2))
	t.Log(client.GetBit(ctx, key3, 3))
	t.Log(client.GetBit(ctx, key3, 4))
	t.Log(client.GetBit(ctx, key3, 5))
	t.Log(client.GetBit(ctx, key3, 6))
	t.Log(client.GetBit(ctx, key3, 7))
}

// 返回字符串中 指定位置1或0的第一个位置
func TestPos(t *testing.T) {
	key := "test:bit_pos:key"
	defer func() {
		client.Del(ctx, key)
	}()
	// 设置第1位为1
	client.SetBit(ctx, key, 0, 1)
	// 设置第2位为1
	client.SetBit(ctx, key, 1, 1)
	// 获取第一个是 0的位置 偏移量 0字节
	t.Log(client.BitPos(ctx, key, 0))
	// 获取第一个是 1的位置 偏移量 0字节
	t.Log(client.BitPos(ctx, key, 1, 0))
	// 获取第一个是 1的位置 偏移量 2字节
	t.Log(client.BitPos(ctx, key, 1, 2))
	// 设置第16位为1
	client.SetBit(ctx, key, 15, 1)
	// 设置第18位为1
	client.SetBit(ctx, key, 17, 1)
	// 获取第一个是 1的位置 偏移量 2字节
	t.Log(client.BitPos(ctx, key, 1, 2))
}

func TestSetBit(t *testing.T) {
	key := "test:set_bit:key"
	client.SetBit(ctx, key, 2, 1)
	client.SetBit(ctx, key, 3, 1)
	client.SetBit(ctx, key, 5, 1)
	client.SetBit(ctx, key, 10, 1)
	client.SetBit(ctx, key, 11, 1)
	client.SetBit(ctx, key, 14, 1)
	t.Log(client.Get(ctx, key))
}
