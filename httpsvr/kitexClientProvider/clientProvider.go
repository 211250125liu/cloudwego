package kitexClientProvider

import (
	"log"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/loadbalance/lbcache"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/njuer/course/cloudwego/httpsvr/idlProvider"
)

// clients is a map of client name to client instance.
// update this map
// provide the client instance to the router
var clients = make(map[string]genericclient.Client)

var etcdResolver discovery.Resolver

func Init() {
	// init etcd resolver
	tEtcdResolver, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	etcdResolver = tEtcdResolver
	// make a Coroutine to update clients from time to time
	go func() {
		for {
			// update clients
			for k := range clients {
				clients[k] = getClientWithEtcd(k)
			}
			time.Sleep(120 * time.Second)
		}
	}()
}

// GetClient returns the client instance by name.
func GetClient(serviceName string) genericclient.Client {
	thisClient, ok := clients[serviceName]
	if !ok {
		thisClient = getClientWithEtcd(serviceName)
		clients[serviceName] = thisClient
		return thisClient
	} else {
		return thisClient
	}
}

// make sure serviceName equals to callServiceName
func getClientWithEtcd(serviceName string) genericclient.Client {
	var opts = []client.Option{
		client.WithResolver(etcdResolver),
		client.WithLoadBalancer(loadbalance.NewWeightedRandomBalancer(), &lbcache.Options{RefreshInterval: 60 * time.Second}),
	}
	var provider = idlProvider.GetIdlByServiceName(serviceName)
	var g, _ = generic.HTTPThriftGeneric(provider)
	var cli, err = genericclient.NewClient(serviceName, g, opts...)
	if err != nil {
		log.Fatal(err)
	}
	return cli
}
