package cache

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
)

func TestFileCacheSet(t *testing.T) {
	fc := New("file")
	re, err := fc.Set("123", "ddd", 0)
	fmt.Println(re)
	if err != nil {
		fmt.Println(err.Error())
		t.Fatalf("写入失败")
	}
}

func TestFileCacheGet(t *testing.T) {
	fc := New("file")
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

func TestRedisCacheSet(t *testing.T) {
	rc := NewRedis(&redis.Options{
		Addr:     "192.168.1.168:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	re, err := rc.Set("123", "ddd", 0)
	fmt.Println(re)
	if err != nil {
		fmt.Println(err.Error())
		t.Fatalf("写入失败")
	}
}

func TestRedisCacheGet(t *testing.T) {
	rc := NewRedis(&redis.Options{
		Addr:     "192.168.1.168:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := rc.Set("123", "451156", 300)
	if err != nil {
		t.Fatalf("写入失败")
	}
	res := ""
	err = rc.Get("123", &res)
	if err != nil {
		t.Fatalf("读取失败")
	}
	fmt.Println("res", res)
}
