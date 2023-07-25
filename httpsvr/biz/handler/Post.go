package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/adaptor"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/generic"
	kitexClientProvider "github.com/njuer/course/cloudwego/httpsvr/kitexClientProvider"
	"github.com/njuer/course/cloudwego/httpsvr/routing"
)

// Post .
func Post(ctx context.Context, c *app.RequestContext) {
	//postInfo := c.Param("postInfo")
	//splitChar := "/"
	//field := strings.Split(postInfo, splitChar)
	//if len(field) != 2 {
	//	panic("You need to specify the service name and query")
	//	c.JSON(200, "error")
	//	return
	//}
	//serviceName := field[0]
	serviceName := routing.GetServiceName(c, "postInfo")
	//for str := range field {
	//	fmt.Println(str)
	//}
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
