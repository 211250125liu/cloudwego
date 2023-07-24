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

// Post .
func Post(ctx context.Context, c *app.RequestContext) {
	postInfo := c.Param("postInfo")
	splitChar := "/"
	field := strings.Split(postInfo, splitChar)
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
	resp, err := client.GenericCall(ctx, serviceName, customReq)
	if err != nil {
		panic(err)
	}

	realResp := resp.(*generic.HTTPResponse)
	c.JSON(consts.StatusOK, realResp.Body)
}
