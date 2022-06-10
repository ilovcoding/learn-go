package gktime_getting_started

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(1 * time.Second)
	return fmt.Sprintf("The result is %d", id)
}
func FirstResponse() string {
	numOfRunner := 10
	// 	ch := make(chan string)
	// 下面的写法也不能防止协程泄露
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			res := runTask(i)
			ch <- res
		}(i)
	}
	return <-ch
}

func TestOne(t *testing.T) {
	t.Log("Before", runtime.NumGoroutine())
	t.Log(FirstResponse())
	time.Sleep(time.Second)
	t.Log("After", runtime.NumGoroutine())
}
