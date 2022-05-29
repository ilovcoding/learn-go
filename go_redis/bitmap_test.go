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
