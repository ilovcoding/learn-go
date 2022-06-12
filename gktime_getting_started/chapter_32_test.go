package gktime_getting_started

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func AllResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			res := runTask(i)
			ch <- res
		}(i)
	}
	for i := 0; i < numOfRunner; i++ {
		v, _ := <-ch
		fmt.Println(v)
	}
	return "all response"
}

func TestAllResponse(t *testing.T) {
	t.Log("Before", runtime.NumGoroutine())
	t.Log(AllResponse())
	time.Sleep(time.Second)
	t.Log("After", runtime.NumGoroutine())
}
