package Gwen

//go:generate swag init -g cmd/apimain.go --output docs/api --instanceName api --exclude http/controller/admin
//go:generate go build -o GwenApi cmd/apimain.go
