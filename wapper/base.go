package wapper

import (
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/metadata"
	"go-micro-service/service/client"
)

func NewLogWrapper(c client.Client) client.Client {
	return &LogWrapper{c}
}

func NewTimeoutCatchWrapper(c client.Client) client.Client {
	return &TimeoutCatchWrapper{c}
}

type LogWrapper struct {
	client.Client
}

type TimeoutCatchWrapper struct {
	client.Client
}

func (logWrapper *LogWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {

	md, _ := metadata.FromContext(ctx)
	fmt.Printf("[Info] ctx: %v service: %s method: %s\n", md, req.Service(), req.Endpoint())
	return logWrapper.Client.Call(ctx, req, rsp)
}

func (timeoutCatchWrapper *TimeoutCatchWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	commandName := req.Service() + "." + req.Endpoint()
	commandConfig := hystrix.CommandConfig{
		Timeout: 1000,
	}
	hystrix.ConfigureCommand(commandName, commandConfig)
	return hystrix.Do(commandName, func() error {
		md, _ := metadata.FromContext(ctx)
		fmt.Printf("[Info] TimeoutOutCatch ctx: %v service: %s method: %s\n", md, req.Service(), req.Endpoint())
		return timeoutCatchWrapper.Client.Call(ctx, req, rsp)
	}, func(err error) error {
		service.GetDataDefault(rsp)
		return nil
	})
}
