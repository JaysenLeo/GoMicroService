package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	Models "go-micro-service/models"
	wepApp "go-micro-service/web"
)

type LogWrapper struct {
	client.Client
}

func (logWrapper *LogWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {

	md, _ := metadata.FromContext(ctx)
	fmt.Printf("[Log Wrapper] ctx: %v service: %s method: %s\n", md, req.Service(), req.Endpoint())
	return logWrapper.Client.Call(ctx, req, rsp)
}

func NewLogWrapper(c client.Client) client.Client {
	return &LogWrapper{c}
}

func main() {
	server := micro.NewService(
		micro.Name("user.client"),
		micro.Registry(etcd.NewRegistry(
			registry.Addrs("127.0.0.1:2379"),
		)),
		micro.WrapClient(NewLogWrapper),
	)
	userServiceClient := Models.NewUserCommonService("user.server", server.Client())

	httpServer := web.NewService(
		web.Name("http.server"),
		web.Address("127.0.0.1:9000"),
		web.Handler(wepApp.NewRouter(userServiceClient)))
	httpServer.Run()

}
