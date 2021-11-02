package global

import (
	"Gwen/config"
	"Gwen/lib/cache"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	LOGGER *logrus.Logger
	Config config.Config
	Viper  *viper.Viper
	Redis  *redis.Client
	Cache  cache.Handler
)
