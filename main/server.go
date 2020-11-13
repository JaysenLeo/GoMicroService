package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	Models "go-micro-service/models"
	server2 "go-micro-service/service/server"
)

func main() {
	reg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	server := micro.NewService(
		micro.Name("user.server"),
		micro.Registry(reg),
	)
	server.Init()
	Models.RegisterUserListServiceHandler(server.Server(), new(server2.UserService))
	server.Run()
}
