package single1

import (
	"fmt"
	"testing"
	"time"
)

type Tools struct {
	value string
}

var instance *Tools

func GetInstance() *Tools {
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
			inst := GetInstance()
			fmt.Printf("%p \n", inst)
		}()
	}
	time.Sleep(time.Second * 2)
}
