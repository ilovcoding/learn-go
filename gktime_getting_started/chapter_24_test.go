package gktime_getting_started

import (
	"fmt"
	"testing"
)

func multipleResFun(a int, b int) (int, int) {
	return a + b, a * b
}

func TestMultipleRes(t *testing.T) {
	res1, res2 := multipleResFun(2, 5)
	t.Log(res1, res2)
}

func intFun(params int) {
	fmt.Printf("params 地址: %p\n", &params)
}

func sliceFun1(params []int) {
	fmt.Printf("params 的地址是 %p\n", &params)
}

func TestFunParams(t *testing.T) {
	a := 10
	fmt.Printf("a的地址是 %p\n", &a)
	intFun(a)
	slice1 := make([]int, 1)
	fmt.Printf("slice的地址是 %p\n", &slice1)
	sliceFun1(slice1)
	t.Log(&slice1)
}

func fun1(a int, b int) int {
	return a + b
}

func TestFuncAsValue(t *testing.T) {
	funcAdd := fun1
	t.Log(funcAdd(1, 10))
}
