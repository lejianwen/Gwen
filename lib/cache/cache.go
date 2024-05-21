package cache

import (
	"github.com/go-redis/redis/v8"
	"sync"
)

type Handler interface {
	Get(key string, value interface{}) error
	Set(key string, value interface{}, exp int) (bool, error)
	Gc() error
	EncodeValue(value interface{}) (string, error)
	DecodeValue(value string, rtv interface{}) error
}

// MaxTimeOut 最大超时时间

const (
	TypeRedis  = "redis"
	TypeFile   = "file"
	MaxTimeOut = 365 * 24 * 3600
)

func New(typ string) Handler {
	var cache Handler
	switch typ {
	case TypeFile:
		cache = NewFileCache()
	case TypeRedis:
		cache = new(RedisCache)
	default:
		cache = NewFileCache()
	}
	return cache
}

func NewRedis(conf *redis.Options) *RedisCache {
	cache := RedisCacheInit(conf)
	return cache
}

func NewFileCache() *FileCache {
	return &FileCache{
		locks: make(map[string]*sync.Mutex),
	}
}
