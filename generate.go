package Gwen

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate swag init -g cmd/main.go
//go:generate swag init -g cmd/main_api.go
//go:generate go run cmd/main.go
