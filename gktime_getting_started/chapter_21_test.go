package gktime_getting_started

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

type Test struct {
}

func TestStringType(t *testing.T) {
	pointer := &Test{}
	t.Log(reflect.TypeOf(pointer))
	str := "str1"
	t.Log(reflect.TypeOf(str))
	t.Log(len(str))
	s := "\xe4\xb8\xA5"
	t.Log(s)
	s = "中"
	t.Log(len(s))
	c := []rune(s)
	t.Log(len(c))
	t.Logf("中 Unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "测试中...𠮷𠮷𠮷"
	for _, c := range s {
		t.Logf("%[1]c %[1]x", c)
	}
}

func TestStringFun(t *testing.T) {
	str := "A,B,C"
	strSlice := strings.Split(str, ",")
	t.Log(strSlice)
	t.Log(strings.Join(strSlice, "_"))
	s := strconv.Itoa(10)
	t.Log("s" + s)
	if num, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + num)
	}
}
