package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	Models "go-micro-service/models"
	"go-micro-service/service"
)

func main() {
	reg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	server := micro.NewService(
		micro.Name("user.server"),
		micro.Registry(reg),
	)
	Models.RegisterUserCommonServiceHandler(server.Server(), new(service.UserService))
	server.Init()
	server.Run()
}
