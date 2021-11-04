package cache

import "github.com/go-redis/redis/v8"

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
	case "file":
		cache = new(FileCache)
	case "redis":
		cache = new(RedisCache)
	default:
		cache = new(FileCache)
	}
	return cache
}

func NewRedis(conf *redis.Options) *RedisCache {
	cache := RedisCacheInit(conf)
	return cache
}

func NewFile() *FileCache {
	return &FileCache{}
}
