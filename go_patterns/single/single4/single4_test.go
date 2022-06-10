package single4

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

func GetInstance() *Tools {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			time.Sleep(time.Second)
			instance = &Tools{
				value: "",
			}
		}
	}
	return instance
}

func TestClient(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			inst := GetInstance()
			fmt.Printf("%p \n", inst)
		}()
	}
	time.Sleep(time.Second * 5)
}
