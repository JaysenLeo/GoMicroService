package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	Models "go-micro-service/models"
	"log"
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
	userServiceClient := Models.NewUserListService("user.server", server.Client())

	resp, e := userServiceClient.GetUserList(context.Background(), &Models.UsersRequest{Size: 5})
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(resp)
	for r := range resp.Data {
		fmt.Println(resp.Data[r].UserID, resp.Data[r].Name)
	}
}
