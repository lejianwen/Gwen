package redis

import (
	"context"
	"fmt"
	"testing"
)

var ctx = context.Background()

func TestNew(t *testing.T) {
	rdb := New("192.168.1.168:6379", "", 0)
	err := rdb.Ping(ctx).Err()
	if err != nil {
		fmt.Println(err)
		t.Fatalf("链接失败")
	}
}

func TestGet(t *testing.T) {
	rdb := New("192.168.1.168:6379", "", 0)
	val, err := rdb.Get(ctx, "test").Result()
	if err != nil {
		fmt.Println(err)
		t.Fatalf("获取失败")
	}
	fmt.Println(val)
}
