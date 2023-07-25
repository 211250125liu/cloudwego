
# cloudwego 使用说明文档  
### 目录结构  
├── Dev.md  
├── Readme.md # 文档所在位置  
├── httpsvr # http端  
│ ├── biz  
│ │ ├── handler  
│ │ │ ├── Get.go # 处理get请求  
│ │ │ ├── Post.go # 处理post请求  
│ │ │ └── ping.go  
│ │ └── router  
│ │ └── register.go  
│ ├── build.sh  
│ ├── go.mod  
│ ├── go.sum  
│ ├── idlProvider #idl文件管理  
│ │ └── provider.go  
│ ├── kitexClientProvider #获取kitex客户端  
│ │ └── clientProvider.go  
│ ├── main.go  
│ ├── router.go  
│ ├── router_gen.go  
│ ├── routing #路由层  
│ │ └── routing.go  
│ └── script  
│ └── bootstrap.sh  
├── idl #idl文件  
│ ├── student.thrift  
│ └── teacher.thrift  
└── rpcsvr  
├── studentservice  
│ ├── build.sh  
│ ├── go.mod  
│ ├── go.sum  
│ ├── handler.go  
│ ├── info.db  
│ ├── kitex_gen  
│ │ └── student  
│ │ ├── k-consts.go  
│ │ ├── k-student.go  
│ │ ├── student.go  
│ │ └── studentservice  
│ │ ├── client.go  
│ │ ├── invoker.go  
│ │ ├── server.go  
│ │ └── studentservice.go  
│ ├── kitex_info.yaml  
│ ├── main.go  
│ ├── output  
│ │ ├── bin  
│ │ │ └── studentservice  
│ │ └── bootstrap.sh  
│ ├── script  
│ │ └── bootstrap.sh  
│ └── test  
│ └── main_test.go  
├── teacherservice  
│ ├── build.sh  
│ ├── go.mod  
│ ├── go.sum  
│ ├── handler.go  
│ ├── info.db  
│ ├── kitex_gen  
│ │ └── teacher  
│ │ ├── k-consts.go  
│ │ ├── k-teacher.go  
│ │ ├── teacher.go  
│ │ └── teacherservice  
│ │ ├── client.go  
│ │ ├── invoker.go  
│ │ ├── server.go  
│ │ └── teacherservice.go  
│ ├── kitex_info.yaml  
│ ├── main.go  
│ ├── output  
│ │ ├── bin  
│ │ │ └── teacherservice  
│ │ └── bootstrap.sh  
│ ├── script  
│ │ └── bootstrap.sh  
│ └── test  
│ └── main_test.go  
└── teacherservicegender #添加gender字段后的teacher服务  
├── build.sh  
├── go.mod  
├── go.sum  
├── handler.go  
├── info.db  
├── kitex_gen  
│ └── teacher  
│ ├── k-consts.go  
│ ├── k-teacher.go  
│ ├── teacher.go  
│ └── teacherservice  
│ ├── client.go  
│ ├── invoker.go  
│ ├── server.go  
│ └── teacherservice.go  
├── kitex_info.yaml  
├── main.go  
├── output  
│ ├── bin  
│ │ └── teacherservice  
│ └── bootstrap.sh  
└── script  
└── bootstrap.sh  
  
### 部署说明  
1. 创建目录idl，编写idl文件  
2. 创建http_server  
mkdir httpsvr  
cd httpsvr  
3. 生成hertz_server脚手架  
hz new -mod github.com /<your_name>/httpsvr -idl ../idl/student.thrift  
4. 生成 kitex client 代码（注意 -module 的值和上面 -mod 的值应当一样）  
kitex -module github.com/<your_name>/httpsvr ../idl/student.thrift  
5. 创建rpc_server  
mkdir rpcsvr  
cd rpcsvr  
6. 生成 kitex server 脚手架  
kitex -module github.com/<your_name>/r  
pcsvr -service student-server ../student.thrift  
修改 main.go 监听与 hertz 不同的端口，  
例如 8889：  
addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8889")  
svr:=demo.NewServer(new(StudentServiceImpl), server.WithServiceAddr(addr))  
7. 运行etcd后即可运行http端，rpc端  
  
### 代码接口说明  
  
1. ![](https://box.nju.edu.cn/f/2568d828076e4731b779/?dl=1)  
http启动时，初始化kitex client provide，进行服务发现，初始化idlProvider，读取idl文件内容，并每隔一段时间进行一次idl更新  
  
2. 在go文件httpsvr/biz/handler/Get.go和httpsvr/biz/handler/Post.go处，调用路由层的GetServiceName方法获得服务名称后，调用kitexClientProvider的GetClient方法获得客户端，后进行泛化调用  
![](https://box.nju.edu.cn/f/1238413d71814d23a79f/?dl=1)  
  
3. idl热更新实现于teacherservice，可以增加一个string类型的gender字段，rpc端由teacherservicegender提供  
  
  
  
# 测试方案说明  
1. 在每个rpc服务下设置test，由http端访问rpc服务得到的数据与实际应得数据进行对比  
2. 同时进行benchmark测试，测试程序的性能  
3. 使用Apache Benchmark进行压测，测试http服务器端的性能  
  
# 性能测试数据  
markbench测试  
![](https://box.nju.edu.cn/f/32cd4926c2d341868ee8/?dl=1)  
  
ab压测  
并发线程数10  
  ![](https://box.nju.edu.cn/f/8e70c39097274fc381c2/?dl=1)
并发线程数100  
![](https://box.nju.edu.cn/f/6064da237af1405285e9/?dl=1)  
# 优化方案说明  
1. 在idl更新时，使用updateIDL方法更新，而不是每个请求生成一个client
![](https://box.nju.edu.cn/f/75eb682f18e0479da3e6/?dl=1)  
2. 使用map缓存客户端，而不是每个请求生成一个client
![](https://box.nju.edu.cn/f/bbf8be78988943159af8/?dl=1)
3. 使用goroutine，减少内存消耗

# 优化后性能数据  
markbench测试
![](https://box.nju.edu.cn/f/35340b81c8d54d57b695/?dl=1)

ab压测  
并发线程数10  
  ![](https://img1.imgtp.com/2023/07/25/Qdjs9dSh.png)
并发线程数100  
![](https://img1.imgtp.com/2023/07/25/fttZb4dN.png)
