## 该项目是自学go语言后，为加深理解实现的一个框架

### 使用的mod参见 go.mod

### 项目目录如下

~~~
app
├── conf                            配置文件
│     └── config.yaml
├── config                          配置相关结构体
│     ├── config.go
│     ├── gin.go
│     ├── gorm.go
│     ├── logger.go
│     └── redis.go
├── global                          全局变量
│     └── global.go
├── go.mod                  
├── go.sum
├── http                            http相关
│     ├── controller                控制器,下面每个子目录对应一个模块
│     │     ├── admin               admin控制器
│     │     │     ├── admin.go
│     │     │     └── login.go
│     │     └── response.go
│     ├── http.go
│     ├── middleware
│     │     ├── admin.go
│     │     └── logger.go
│     ├── router                    路由,下面每个子目录对应一个模块
│     │     ├── admin               admin路由，下面每个路由文件对应一个controller
│     │     │     ├── admin.go
│     │     │     └── login.go
│     │     ├── admin.go            admin路由注册
│     │     ├── api.go              api路由注册
│     │     └── router.go
│     ├── run.go                    server启动,非win系统下
│     └── run_win.go                server启动,win系统下
├── lib                             核心包
│     ├── cache                     缓存
│     │     ├── cache.go
│     │     ├── cache_test.go
│     │     ├── file.go
│     │     ├── file_lock.go
│     │     ├── file_lock_win.go
│     │     ├── file_test.go
│     │     ├── redis.go
│     │     └── redis_test.go
│     ├── logger                    日志
│     │     └── logger.go
│     └── redis                     redis
│         ├── redis.go
│         └── redis_test.go
├── main.go
├── model                           gorm模型
│     ├── admin.go
│     ├── adminRole.go
│     └── model.go                  gorm初始化
├── runtime                         运行日志，缓存目录
│     └── log.txt
├── service                         service层，dao操作放到service中，不单独再写dao层
│     ├── admin.go
│     └── service.go
└── utils                           其他工具
    └── response.go


~~~

# 项目仅供学习，随时可能改动