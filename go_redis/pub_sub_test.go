package go_redis

import (
	"fmt"
	"testing"
)

// 发布订阅

func TestPubSub(t *testing.T) {
	key := "test:pub:sub:key1"
	// 只有消费者先监听才能收到消息
	cSub := client.Subscribe(ctx, key)
	t.Log(cSub)
	go func() {
		for i := 0; i < 10; i++ {
			str := fmt.Sprintf("data %d", i)
			client.Publish(ctx, key, str)
			//t.Log(res)
		}
	}()
	for {
		msg, _ := cSub.ReceiveMessage(ctx)
		t.Log(msg)
	}
}
