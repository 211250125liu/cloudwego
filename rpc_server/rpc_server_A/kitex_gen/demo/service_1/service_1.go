// Code generated by Kitex v0.6.1. DO NOT EDIT.

package service_1

import (
	"context"
	demo "github.com/211250125liu/rpc_server/kitex_gen/demo"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return service_1ServiceInfo
}

var service_1ServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "service_1"
	handlerType := (*demo.Service_1)(nil)
	methods := map[string]kitex.MethodInfo{
		"getMessage": kitex.NewMethodInfo(getMessageHandler, newService_1GetMessageArgs, newService_1GetMessageResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "demo",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.6.1",
		Extra:           extra,
	}
	return svcInfo
}

func getMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*demo.Service_1GetMessageArgs)
	realResult := result.(*demo.Service_1GetMessageResult)
	success, err := handler.(demo.Service_1).GetMessage(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newService_1GetMessageArgs() interface{} {
	return demo.NewService_1GetMessageArgs()
}

func newService_1GetMessageResult() interface{} {
	return demo.NewService_1GetMessageResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetMessage(ctx context.Context, req *demo.Request) (r *demo.Response, err error) {
	var _args demo.Service_1GetMessageArgs
	_args.Req = req
	var _result demo.Service_1GetMessageResult
	if err = p.c.Call(ctx, "getMessage", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
