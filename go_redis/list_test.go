package go_redis

import (
	"testing"
	"time"
)

// 当List 里面有数据时候 BLMove 和 LMove 操作效果一样，当list里面没有数据的时候BLMove 会阻塞住当前代码运行，直到队列里面有数据。
// 其他操作和 LMove 相同，具体可以参考 TestLMove 函数
func TestBLMove(t *testing.T) {
	key := "test:BLMove:k1"
	key2 := "test:BLMove:k2"
	defer func() {
		client.Del(ctx, key, key2)
	}()
	client.RPush(ctx, key, "one")
	client.RPush(ctx, key, "two")
	client.RPush(ctx, key, "three")
	res := client.BLMove(ctx, key, key2, "RIGHT", "LEFT", 0)
	t.Log(res)
	res = client.BLMove(ctx, key, key2, "LEFT", "RIGHT", 0)
	t.Log(res)
	res = client.BLMove(ctx, key, key2, "LEFT", "RIGHT", 0)
	t.Log(res)
	// 超时失败
	res = client.BLMove(ctx, key, key2, "LEFT", "RIGHT", time.Second*5)
	t.Log(res)
	go func() {
		time.Sleep(time.Second * 2)
		client.RPush(ctx, key, "four")
	}()
	res = client.BLMove(ctx, key, key2, "LEFT", "RIGHT", 0)
	t.Log(res)
}

/**
Available since: 6.2.0
Time complexity: O(1)
LMOVE source destination LEFT | RIGHT LEFT | RIGHT
从 6.2.0 开始支持此命令 从 source list 头或尾取出一个数，添加到 destination 的头或尾部
*/
func TestLMove(t *testing.T) {
	key := "test:LMove:k1"
	key2 := "test:LMove:k2"

	defer func() {
		client.Del(ctx, key, key2)
	}()
	client.RPush(ctx, key, "one")
	client.RPush(ctx, key, "two")
	client.RPush(ctx, key, "three")
	// key 中的 three 移动到 key2 队头
	client.LMove(ctx, key, key2, "RIGHT", "LEFT")
	// key 中的 one 移动到 key2 队尾
	client.LMove(ctx, key, key2, "LEFT", "RIGHT")
	// two
	res := client.LRange(ctx, key, 0, -1)
	t.Log(res)
	// three one
	res = client.LRange(ctx, key2, 0, -1)
	t.Log(res)
	// 将 ke2 内部的值进行反转
	client.LMove(ctx, key2, key2, "LEFT", "RIGHT")
	res = client.LRange(ctx, key2, 0, -1)
	t.Log(res)
}
