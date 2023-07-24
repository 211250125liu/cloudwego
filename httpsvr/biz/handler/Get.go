package handler

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/adaptor"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/generic"
	kitexClientProvider "github.com/njuer/course/cloudwego/httpsvr/kitexClientProvider"
)

// Get .
func Get(ctx context.Context, c *app.RequestContext) {
	// url
	// https://cn.bing.com/search?q=1
	// /get/servicename/query?id=1

	getInfo := c.Param("getInfo")
	splitChar := "/"
	field := strings.Split(getInfo, splitChar)
	if len(field) != 2 {
		panic("You need to specify the service name and query")
		c.JSON(200, "error")
		return
	}
	serviceName := field[0]
	for str := range field {
		fmt.Println(str)
	}

	client := kitexClientProvider.GetClient(serviceName)
	httpReq, _ := adaptor.GetCompatRequest(c.GetRequest())
	customReq, _ := generic.FromHTTPRequest(httpReq)
	resp, _ := client.GenericCall(ctx, serviceName, customReq)

	realResp := resp.(*generic.HTTPResponse)
	c.JSON(consts.StatusOK, realResp.Body)
}
