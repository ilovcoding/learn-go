package go_redis

import (
	"github.com/go-redis/redis/v8"
	"testing"
)

const zSetKey = "test:zSet:key"

func TestZAdd(t *testing.T) {
	client.ZAdd(ctx, zSetKey, &redis.Z{
		Member: "one",
		Score:  1,
	}, &redis.Z{
		Member: "uno",
		Score:  1,
	}, &redis.Z{
		Member: "z_two",
		Score:  2,
	}, &redis.Z{
		Member: "two",
		Score:  2,
	}, &redis.Z{
		Member: "three",
		Score:  3,
	})
	res := client.ZRange(ctx, zSetKey, 0, -1)
	t.Log(res)
}

func TestZIncrBy(t *testing.T) {
	client.ZAdd(ctx, zSetKey, &redis.Z{
		Member: "one",
		Score:  1,
	}, &redis.Z{
		Member: "two",
		Score:  2,
	})
	client.ZIncrBy(ctx, zSetKey, 2, "one")
	res := client.ZRangeWithScores(ctx, zSetKey, 0, -1)
	t.Log(res)
}

func TestZScore(t *testing.T) {
	res := client.ZScore(ctx, zSetKey, "one")
	t.Log(res)
}
