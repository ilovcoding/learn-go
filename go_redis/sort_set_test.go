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

func TestZRank(t *testing.T) {
	res := client.ZRangeWithScores(ctx, zSetKey, 0, -1)
	t.Log(res)
	res2 := client.ZRank(ctx, zSetKey, "one")
	t.Log(res2)
	res2 = client.ZRank(ctx, zSetKey, "uno")
	t.Log(res2)
	res2 = client.ZRank(ctx, zSetKey, "two")
	t.Log(res2)
	res2 = client.ZRank(ctx, zSetKey, "z_two")
	t.Log(res2)
	res2 = client.ZRank(ctx, zSetKey, "three")
	t.Log(res2)
}

func TestZScore(t *testing.T) {
	res := client.ZScore(ctx, zSetKey, "one")
	t.Log(res)
}
