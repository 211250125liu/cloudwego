package main

import (
	"context"
	"fmt"
	demo "github.com/211250125liu/rpc_server/kitex_gen/demo/service_1"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	ruleBasedResolver "github.com/kitex-contrib/resolver-rule-based"
	"log"
	"net"
	"strconv"
)

var (
	etcdAddr     = "127.0.0.1:2379"
	serviceName  = "rpc_server_A"
	tagKey       = "k"
	tagValues    = []string{"v1", "v2"}
	instanceTags = []map[string]string{
		{tagKey: tagValues[0]},
		{tagKey: tagValues[1]},
	}
)

func resolve() {
	// use etcd resolver
	etcdResolver, err := etcd.NewEtcdResolver([]string{etcdAddr})
	if err != nil {
		panic(err)
	}
	filterFunc := func(ctx context.Context, instance []discovery.Instance) []discovery.Instance {
		var res []discovery.Instance
		for _, ins := range instance {
			if v, ok := ins.Tag(tagKey); ok && v == tagValues[0] {
				// only match tag with {tagKey: tagValues[0]}
				res = append(res, ins)
			}
		}
		return res
	}
	// Construct the filterRule
	filterRule := &ruleBasedResolver.FilterRule{Name: "rule-name", Funcs: []ruleBasedResolver.FilterFunc{filterFunc}}
	// build rule based resolver
	rbr := ruleBasedResolver.NewRuleBasedResolver(etcdResolver, filterRule)

	// service discovery
	ctx := context.Background()
	ei := rpcinfo.NewEndpointInfo(serviceName, "getMessage", nil, nil)
	desc := rbr.Target(ctx, ei)
	res, err := rbr.Resolve(ctx, desc)
	if err != nil {
		panic(err)
	}

	// the instance should match the filter rule
	v, _ := res.Instances[0].Tag(tagKey)
	fmt.Println(fmt.Sprintf(
		"[Resolver]: get instance with tag, [%s:%s]", tagKey, v))
}

func main() {
	// instances
	var instances []*registry.Info
	for i := 0; i < len(instanceTags); i++ {
		addr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf("127.0.0.1:%s", strconv.Itoa(8888+i)))
		instances = append(instances, &registry.Info{
			ServiceName: serviceName,
			Addr:        addr,
			Tags:        instanceTags[i],
		})
	}

	//服务注册
	r, err := etcd.NewEtcdRegistry([]string{etcdAddr})
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(instances); i++ {
		err := r.Register(instances[i])
		if err != nil {
			panic(err)
		}
	}
	defer func() {
		for i := 0; i < len(instances); i++ {
			err := r.Deregister(instances[i])
			if err != nil {
				return
			}
		}
	}()
	resolve()

	//ebi := &rpcinfo.EndpointBasicInfo{
	//	ServiceName: "rpc_server_A",
	//	Tags:        make(map[string]string),
	//}
	//ebi.Tags["rpc_server_A"] = "rpc_server_A"
	//
	//addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8889")
	//
	//svr := demo.NewServer(new(Service_1Impl), server.WithServerBasicInfo(ebi),
	//	server.WithRegistry(r), server.WithServiceAddr(addr),
	//	server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
	//		ServiceName: "rpc_server_A",
	//		Method:      "GetMessage",
	//		Tags:        nil,
	//	}))
	//err = svr.Run()
	//if err != nil {
	//	log.Println(err.Error())
	//}

	//r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	//if err != nil {
	//	log.Fatal(err)
	//}

	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8889")
	svr := demo.NewServer(new(Service_1Impl), server.WithRegistry(r), server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "rpc_server_A",
			Method:      "GetMessage",
			Tags:        nil,
		}))

	//svr := demo.NewServer(new(Service_1Impl), server.WithServiceAddr(addr))

}
