package bootstrap

import (
	"Gwen/config"
	"Gwen/global"
	"Gwen/http"
	"Gwen/lib/cache"
	"Gwen/lib/logger"
	"Gwen/lib/redis"
	"Gwen/model"
	"fmt"
	redis2 "github.com/go-redis/redis/v8"
)

func Init() {
	global.Viper = config.Init(&global.Config, func() {
		fmt.Println(global.Config)
	})

	global.LOGGER = logger.New(logger.Config{
		Path:         global.Config.Logger.Path,
		Level:        global.Config.Logger.Level,
		ReportCaller: global.Config.Logger.ReportCaller,
		Mode:         global.Config.Logger.Mode,
	})

	global.Redis = redis.New(global.Config.Redis.Addr, global.Config.Redis.Password, global.Config.Redis.Db)

	if global.Config.Cache.Type == "file" {
		fc := cache.NewFile()
		fc.SetDir(global.Config.Cache.FileDir)
		global.Cache = fc
	} else if global.Config.Cache.Type == "redis" {
		global.Cache = cache.NewRedis(&redis2.Options{
			Addr:     global.Config.Cache.RedisAddr,
			Password: global.Config.Cache.RedisPwd,
			DB:       global.Config.Cache.RedisDb,
		})
	}

	global.DB = model.Init()
	http.Init()
}
