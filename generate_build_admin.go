package Gwen

//go:generate swag init -g cmd/main.go --output docs/api --instanceName api --exclude http/controller/admin
//go:generate go build -o GwenAdmin cmd/main.go
