package cache

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFileSet(t *testing.T) {
	fc := &FileCache{}
	re, err := fc.Set("123", "ddd", 0)
	fmt.Println(Dir)
	fmt.Println(re)
	if err != nil {
		fmt.Println(err.Error())
		t.Fatalf("写入失败")
	}
}

func TestFileGet(t *testing.T) {
	fc := &FileCache{}
	_, err := fc.Set("123", "45156", 300)
	if err != nil {
		t.Fatalf("写入失败")
	}
	res := ""
	err = fc.Get("123", &res)
	if err != nil {
		t.Fatalf("读取失败")
	}
	fmt.Println("res", res)
}

func TestFileGetJson(t *testing.T) {
	fc := &FileCache{}
	type r struct {
		Aa string `json:"a"`
		B  string `json:"c"`
	}
	old := &r{
		Aa: "ab", B: "cdc",
	}
	_, err := fc.Set("123", old, 300)
	if err != nil {
		t.Fatalf("写入失败")
	}

	res := &r{}
	err2 := fc.Get("123", res)
	if err2 != nil {
		t.Fatalf("读取失败")
	}
	if !reflect.DeepEqual(res, old) {
		t.Fatalf("读取错误")
	}
	fmt.Println(res, res.Aa)
}

func BenchmarkSet(b *testing.B) {
	fc := &FileCache{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fc.Set("123", "{dsv}", 1000)
	}
}

func BenchmarkGet(b *testing.B) {
	fc := &FileCache{}
	b.ResetTimer()
	v := ""
	for i := 0; i < b.N; i++ {
		fc.Get("123", &v)
	}
}
