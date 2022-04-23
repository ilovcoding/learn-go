package single5

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
var once sync.Once

func getInstance() *Tools {
	once.Do(func() {
		time.Sleep(time.Second)
		instance = &Tools{
			value: "",
		}
	})
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
