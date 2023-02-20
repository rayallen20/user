package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	go_micro_service_user "github.com/rayallen20/user/proto/user"
)

func main() {
	userService := micro.NewService(
		micro.Name("go.micro.service.user.client"),
	)

	userService.Init()

	user := go_micro_service_user.NewUserService("go.micro.service.user", userService.Client())

	response, err := user.Login(context.TODO(), &go_micro_service_user.LoginRequest{
		Account:  "admin",
		Password: "NEWnew123!@#",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v\n", response)
}
