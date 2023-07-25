package routing

import (
	"github.com/cloudwego/hertz/pkg/app"
	"strings"
)

func GetServiceName(c *app.RequestContext, key string) string {
	info := c.Param(key)
	splitChar := "/"
	field := strings.Split(info, splitChar)
	if len(field) != 2 {
		c.JSON(200, "error")
		panic("You need to specify the service name and query")
		return ""
	}
	return field[0]
}
