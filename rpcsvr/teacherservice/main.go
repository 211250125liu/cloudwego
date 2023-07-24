package main

import (
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/njuer/course/cloudwego/rpcteasvr/kitex_gen/teacher/teacherservice"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	addr, _ := net.ResolveTCPAddr("tcp", ":9999")
	svr := teacherservice.NewServer(new(TeacherServiceImpl).InitDB(), server.WithRegistry(r),
		server.WithServerBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: "teacher",
			},
		),
		server.WithServiceAddr(addr),
	)
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
