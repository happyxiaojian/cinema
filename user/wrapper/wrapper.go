package wrapper

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
	"time"
)

// handler 包装
func HandlerWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		fmt.Printf("[HandlerWrapper] [%v] server request: %s\n", time.Now(), req.Method())
		return fn(ctx, req, rsp)
	}
}

type clientWrapper struct {
	client.Client
}

func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	fmt.Printf("[wrapper] client request to service: %s method: %s\n", req.Service(), req.Method())
	return c.Client.Call(ctx, req, rsp)
}

// 返回一个包装过的客户端
func LogClientWrap(c client.Client) client.Client {
	return &clientWrapper{c}
}

func CallFuncWrap(callFunc client.CallFunc) client.CallFunc {
	return func(ctx context.Context, address string, req client.Request, rsp interface{}, opts client.CallOptions) error {
		fmt.Println("[wrapper] call func : ", ctx, address, req, rsp, opts)
		return callFunc(ctx, address, req, rsp, opts)
	}
}