## 该项目是自学go语言后，为加深理解实现的一个框架

### 项目分为admin和api

- admin是后台接口,在cmd/main.go中启动
- api是接口系统,在cmd/apimain.go中启动

### 项目使用的框架如下

- gin
- gorm
- viper
- swag
- go-redis
- validator
- endless
- jwt-go
- logrus
- ...

### 自己实现了一些简单的功能

- cache
- lock
- upload

## 使用方式

### 安装依赖

```shell
go mod tidy
```

### 直接运行

```shell
#api
go run cmd/apimain.go 
#或者指定配置文件
go run cmd/apimain.go -c ./conf/config.yaml

#admin
go run cmd/main.go -c ./conf/config.yaml
```

## 使用generate运行

### 1.先设置要编译到的环境

```shell
#linux
go generate generate_env_linux.go

#或者windows
go generate generate_env_win.go
```

### 2.生成文档,然后编译或者直接运行

```shell
#运行api
go generate generate_api.go

#编译api
go generate generate_build_api.go

```

### 3.部署

#### 先上传***conf***目录, ***runtime***目录, ***resources***目录到正式环境

#### 再上传生成好的***GwenApi***或者***GwenAdmin***,并给与运行权限```chmod +x GwenApi```

```shell
#到目录下
./GwenApi
#再不同的环境下，指定不同的配置文件即可，比如pro
./GwenApi -c ./conf/config_pro.yaml
```

### 使用的mod参见 go.mod

### 项目目录如下

~~~
app
├── cmd
│   ├── apimain.go
│   └── main.go
├── conf
│   ├── config.yaml
│   └── jwt_pri.pem   #jwt私钥，自己生成
├── config
│   ├── cache.go
│   ├── config.go
│   ├── gin.go
│   ├── gorm.go
│   ├── jwt.go
│   ├── logger.go
│   ├── oss.go
│   └── redis.go
├── docs
│   ├── admin
│   │   ├── admin_docs.go
│   │   ├── admin_swagger.json
│   │   └── admin_swagger.yaml
│   └── api
│       ├── api_docs.go
│       ├── api_swagger.json
│       └── api_swagger.yaml
├── global
│   └── global.go
├── http
│   ├── controller
│   │   ├── admin
│   │   │   ├── admin.go
│   │   │   ├── admin_role.go
│   │   │   ├── file.go
│   │   │   └── login.go
│   │   └── api
│   │       ├── index.go
│   │       └── user.go
│   ├── middleware
│   │   ├── admin.go
│   │   ├── cors.go
│   │   ├── jwt.go
│   │   └── logger.go
│   ├── request
│   │   ├── admin.go
│   │   ├── admin_role.go
│   │   └── login.go
│   ├── response
│   │   ├── response.go
│   │   └── user.go
│   ├── router
│   │   ├── admin.go
│   │   ├── api.go
│   │   └── router.go
│   ├── http.go
│   ├── run.go
│   └── run_win.go
├── lib
│   ├── cache
│   │   ├── cache.go
│   │   ├── cache_test.go
│   │   ├── file.go
│   │   ├── file_test.go
│   │   ├── memory.go
│   │   ├── memory_test.go
│   │   ├── redis.go
│   │   ├── redis_test.go
│   │   ├── simple_cache.go
│   │   └── simple_cache_test.go
│   ├── jwt
│   │   ├── jwt.go
│   │   └── jwt_test.go
│   ├── lock
│   │   ├── local.go
│   │   ├── local_test.go
│   │   └── lock.go
│   ├── logger
│   │   └── logger.go
│   ├── orm
│   │   └── mysql.go
│   └── upload
│       ├── local.go
│       └── oss.go
├── model
│   ├── custom_types
│   │   ├── auto_json.go
│   │   └── auto_time.go
│   ├── admin.go
│   ├── adminRole.go
│   ├── model.go
│   └── user.go
├── resources
│   └── public
│       └── upload
├── runtime
│   ├── cache
│   └── log.txt
├── service
│   ├── admin.go
│   ├── admin_role.go
│   ├── service.go
│   └── user.go
├── utils
│   └── tools.go
├── generate_admin.go
├── generate_api.go
├── generate_build_api.go
├── generate_env_linux.go
├── generate_env_win.go
├── go.mod
├── go.sum
└── README.md

~~~

# 项目仅供学习，随时可能改动
