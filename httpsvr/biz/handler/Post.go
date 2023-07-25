package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/adaptor"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/njuer/course/cloudwego/httpsvr/kitexClientProvider"
	"github.com/njuer/course/cloudwego/httpsvr/routing"
)

// Post .
func Post(ctx context.Context, c *app.RequestContext) {
	serviceName := routing.GetServiceName(c, "postInfo")

	client := kitexClientProvider.GetClient(serviceName)
	httpReq, err := adaptor.GetCompatRequest(c.GetRequest())
	if err != nil {
		panic("get httpReq failed")
	}
	customReq, err := generic.FromHTTPRequest(httpReq)
	if err != nil {
		panic("get customReq failed")
	}
	resp, err := client.GenericCall(ctx, serviceName, customReq)
	if err != nil {
		panic("generic failed")
	}

	realResp := resp.(*generic.HTTPResponse)
	c.JSON(consts.StatusOK, realResp.Body)
}
