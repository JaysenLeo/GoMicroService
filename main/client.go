package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	Models "go-micro-service/models"
	"go-micro-service/wapper"
	wepApp "go-micro-service/web"
)

func main() {
	server := micro.NewService(
		micro.Name("user.client"),
		micro.Registry(etcd.NewRegistry(
			registry.Addrs("127.0.0.1:2379"),
		)),
		micro.WrapClient(wapper.NewLogWrapper),
		micro.WrapClient(wapper.NewTimeoutCatchWrapper),
	)
	userServiceClient := Models.NewUserCommonService("user.server", server.Client())

	httpServer := web.NewService(
		web.Name("http.server"),
		web.Address("127.0.0.1:9000"),
		web.Handler(wepApp.NewRouter(userServiceClient)))
	httpServer.Run()

}
