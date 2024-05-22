package cache

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFileSet(t *testing.T) {
	fc := NewFileCache()
	re, err := fc.Set("123", "ddd", 0)
	fmt.Println(re)
	if err != nil {
		fmt.Println(err.Error())
		t.Fatalf("写入失败")
	}
}

func TestFileGet(t *testing.T) {
	fc := NewFileCache()
	res := ""
	err := fc.Get("123", &res)
	if err != nil {
		fmt.Println(err.Error())
		t.Fatalf("读取失败")
	}
	fmt.Println("res", res)
}

func TestFileGetJson(t *testing.T) {
	fc := NewFileCache()
	type r struct {
		Aa string `json:"a"`
		B  string `json:"c"`
	}

	res := &r{
		Aa: "ab", B: "cdc",
	}
	err2 := fc.Get("123", res)
	fmt.Println("res", res)
	if err2 != nil {
		t.Fatalf("读取失败" + err2.Error())
	}
}
func TestFileSetGetJson(t *testing.T) {
	fc := NewFileCache()
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
	fmt.Println("res", res)
	if err2 != nil {
		t.Fatalf("读取失败" + err2.Error())
	}
	if !reflect.DeepEqual(res, old) {
		t.Fatalf("读取错误")
	}
	fmt.Println("res", res)
}

func BenchmarkSet(b *testing.B) {
	fc := NewFileCache()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fc.Set("123", "{dsv}", 1000)
	}
}

func BenchmarkGet(b *testing.B) {
	fc := NewFileCache()
	b.ResetTimer()
	v := ""
	for i := 0; i < b.N; i++ {
		fc.Get("123", &v)
	}
}
