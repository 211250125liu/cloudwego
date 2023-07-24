package main

import (
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/njuer/course/cloudwego/rpcstusvr/kitex_gen/student/studentservice"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	addr, _ := net.ResolveTCPAddr("tcp", ":9998")
	svr := studentservice.NewServer(new(StudentServiceImpl).InitDB(), server.WithRegistry(r),
		server.WithServerBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: "student",
			},
		),
		server.WithServiceAddr(addr),
	)
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
