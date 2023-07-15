# cloudwego
首先接受http请求（hertz的部分），然后处理对应的请求，看对应哪一个rpc服务，进行泛化调用，
这一系列的rpc服务由注册中心管理，同时http和rpc交互对应由一系列的thrift协议，
这里要对应那个idl的热更新

## http_server
处理对应的请求，对rpc进行泛化调用

## idl
储存thrift文件，暂时不考虑使用管理平台

## registry
注册中心，etcd的处理可能放这吧

## rpc_server
处理http请求并返回

## test
写最后的集成测试
