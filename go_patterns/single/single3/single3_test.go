package single3

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Tools struct {
	value string
}

var instance *Tools
var mu sync.Mutex

func getInstance() *Tools {
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		time.Sleep(time.Second)
		instance = &Tools{
			value: "",
		}
	}
	return instance
}

func TestClient(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			inst := getInstance()
			fmt.Printf("%p \n", inst)
		}()
	}
	time.Sleep(time.Second * 2)
}
