package single2

import (
	"fmt"
	"testing"
	"time"
)

type Tools struct {
	value string
}

var instance *Tools

func init() {
	fmt.Printf("%p \n", instance)
	instance = &Tools{
		value: "",
	}
}
func GetInstance() *Tools {
	time.Sleep(time.Second)
	return instance
}

func TestClient(t *testing.T) {
	for i := 0; i < 20; i++ {
		go func() {
			inst := GetInstance()
			fmt.Printf("%p \n", inst)
		}()
	}
	time.Sleep(time.Second * 2)
}
