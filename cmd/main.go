package main

import (
	"Gwen/config"
	"Gwen/global"
	"Gwen/http"
	"Gwen/lib/cache"
	"Gwen/lib/lock"
	"Gwen/lib/logger"
	"Gwen/lib/orm"
	"Gwen/lib/upload"
	"Gwen/utils"
	"fmt"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh_Hans_CN"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/go-redis/redis/v8"
	"reflect"
)

// @title 后台管理系统API
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
	})

	//redis
	global.Redis = redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.Addr,
		Password: global.Config.Redis.Password,
		DB:       global.Config.Redis.Db,
	})

	//cache
	if global.Config.Cache.Type == cache.TypeFile {
		fc := cache.NewFileCache()
		fc.SetDir(global.Config.Cache.FileDir)
		global.Cache = fc
	} else if global.Config.Cache.Type == cache.TypeRedis {
		global.Cache = cache.NewRedis(&redis.Options{
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
	InitValidator()

	//oss
	global.Oss = &upload.Oss{}
	utils.CopyStructByJson(&global.Config.Oss, global.Oss)

	//locker
	global.Lock = lock.NewLocal()

	//gin
	http.Init()

}

func InitValidator() {
	validate := validator.New()
	enT := en.New()
	cn := zh_Hans_CN.New()
	uni := ut.New(enT, cn)
	trans, _ := uni.GetTranslator("cn")
	err := zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		//退出
		panic(err)
	}
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

}
