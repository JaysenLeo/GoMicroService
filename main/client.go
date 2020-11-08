package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	Models "go-micro-service/models"
	"log"
)

func main() {
	server := micro.NewService(
		micro.Name("user.client"),
		micro.Registry(etcd.NewRegistry(
			registry.Addrs("127.0.0.1:2379"),
		)),
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
