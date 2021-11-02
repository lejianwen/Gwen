package redis

import (
	"github.com/go-redis/redis/v8"
)

func New(addr, pwd string, db int) (rdb *redis.Client) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd, // no password set
		DB:       db,  // use default DB
	})
	return
}
