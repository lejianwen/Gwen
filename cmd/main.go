package main

import (
	"Gwen/config"
	"Gwen/global"
	"Gwen/http"
	"Gwen/lib/cache"
	"Gwen/lib/logger"
	"Gwen/lib/orm"
	"Gwen/lib/redis"
	"fmt"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh_Hans_CN"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	redis2 "github.com/go-redis/redis/v8"
	"reflect"
)

// @title 供应商后台管理系统API
// @version 1.0
// @description 后台接口
// @basePath /admin-api
// @securityDefinitions.apikey token
// @in header
// @name api-token
func main() {
	//配置解析
	global.Viper = config.Init(&global.Config, func() {
		fmt.Println(global.Config)
	})

	//日志
	global.Logger = logger.New(&logger.Config{
		Path:         global.Config.Logger.Path,
		Level:        global.Config.Logger.Level,
		ReportCaller: global.Config.Logger.ReportCaller,
		Mode:         global.Config.Logger.Mode,
	})

	//redis
	global.Redis = redis.New(global.Config.Redis.Addr, global.Config.Redis.Password, global.Config.Redis.Db)

	//cache
	if global.Config.Cache.Type == cache.TypeFile {
		fc := cache.NewFile()
		fc.SetDir(global.Config.Cache.FileDir)
		global.Cache = fc
	} else if global.Config.Cache.Type == cache.TypeRedis {
		global.Cache = cache.NewRedis(&redis2.Options{
			Addr:     global.Config.Cache.RedisAddr,
			Password: global.Config.Cache.RedisPwd,
			DB:       global.Config.Cache.RedisDb,
		})
	}

	//gorm
	conf := &orm.MysqlConfig{
		Dns: global.Config.Mysql.Dns, MaxIdleConns: global.Config.Mysql.MaxIdleConns, MaxOpenConns: global.Config.Mysql.MaxOpenConns,
	}
	global.DB = orm.NewMysql(conf)

	//validator
	validate := validator.New()
	en := en.New()
	cn := zh_Hans_CN.New()
	uni := ut.New(en, cn)
	trans, _ := uni.GetTranslator("cn")
	zh_translations.RegisterDefaultTranslations(validate, trans)
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		if label == "" {
			return field.Name
		}
		return label
	})
	global.Validator.Validate = validate
	global.Validator.VTrans = trans

	global.Validator.ValidStruct = func(i interface{}) []string {
		err := global.Validator.Validate.Struct(i)
		errList := make([]string, 0, 10)
		if err != nil {
			if _, ok := err.(*validator.InvalidValidationError); ok {
				errList = append(errList, err.Error())
				return errList
			}
			for _, err2 := range err.(validator.ValidationErrors) {
				errList = append(errList, err2.Translate(global.Validator.VTrans))
			}
		}
		return errList
	}
	global.Validator.ValidVar = func(field interface{}, tag string) []string {
		err := global.Validator.Validate.Var(field, tag)
		fmt.Println(err)
		errList := make([]string, 0, 10)
		if err != nil {
			if _, ok := err.(*validator.InvalidValidationError); ok {
				errList = append(errList, err.Error())
				return errList
			}
			for _, err2 := range err.(validator.ValidationErrors) {
				errList = append(errList, err2.Translate(global.Validator.VTrans))
			}
		}
		return errList
	}

	//gin
	http.Init()
}
