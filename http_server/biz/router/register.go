// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	demo "github.com/211250125liu/httpserver/biz/router/demo"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	demo.Register(r)
}
