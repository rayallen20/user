package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/rayallen20/user/handler"

	user "github.com/rayallen20/user/proto/user"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
		// 指定服务端口
		micro.Address("0.0.0.0:8091"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	user.RegisterUserHandler(service.Server(), &handler.User{})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
