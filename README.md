# gofiber-integrate
gofiber intergrate with other framework，gofiber框架整合

## 1.gofiber + swagger 接口文档
> 由于 Fiber 使用了 unsafe 特性，导致其可能与最新的 Go 版本不兼容。Fiber 2.18.0 已经在 Go 1.14 到 1.17 上验证过。Fiber 与 net/http 接口不兼容。这意味着你无法使用 gqlen，go-swagger 或者任何其他属于 net/http 生态的项目。

## 2. gorm 连接mysql 数据库



## 3. go-redis 操作 redis



## 4. gofiber中间jwt实现认证



## 5. gofiber结合go-gomicro 实现微服务

### go语言自带的rpc包

go语言自带了rpc依赖
1. 调用客户端句柄，执行传送参数
2. 调用本地系统内核发送网络消息
3. 消息传递到远程主机
4. 服务器句柄得到消息并取得参数
5. 执行远程过程
6. 执行的过程将结果返回服务器句柄
7. 服务器句柄返回结果，调用远程系统内核
8. 消息传回本地主机
9. 客户端句柄由内核接收消息
10. 客户端接收句柄返回的数据

### GRPC

> Google 开源的高性能，开源和通用的RPC框架，面向移动和HTTP/2设计。基于HTTP/2标准设计，有双向流，流控，头部压缩，单TCP连接上的多复用请求等待。在移动设备上表现更好，更省电和节省空间。
>
> GRPC使用`protocol buffers`的数据格式,可以使用GO,Python,Ruby等来创建客户端



#### protobuf

> google开发的一种与平台语言无关的可拓展序列化数据结构，适合用于不同语言之间的通信，或数据存储

安装protoc

安装protoc-gen-go 插件

编写proto文件  

```protobuf
message Test{
  string name =1;
  repeated int32 weight = 2;
  int32 height = 3;
  string motto = 4;
}
```

编译proto文件

```shell
#将proto文件编译为go语言
protoc --go_out=./  *.proto  #不添加插件的方式
protoc --go_out=plugins=mygrpc:./  *.proto #添加插件的方式
```

#### 编写一个简单grpc的服务

### Consul服务发现

> consul 是HashCorp公司推出的开源工具，用于实现分布式的服务发现与配置。
>
> 服务注册、服务发现、健康检测、存储动态配置

下载consul

 

#### cousul角色

- **client 客户端**：无状态，将http服务和DNS接口请求转发给局域网内的服务器集群
- **server 服务端**，保存配置信息，高可用集群，在局域网内与本地客户端通讯，通过广域网与其他数据中心通讯，每个 数据中心的server数量推荐为3个或5个

  

```sh
consul agent -server -bootstrap-expect 2 -data-dir /data/consul-data -config-dir /etc/consul.d/ -node=n1 -bind=10.211.55.11 -ui  -rejoin -join 10.211.55.11 -client 0.0.0.0

./consul agent -server -bootstrap-expect 2 -data-dir /data/consul-data -node=n2 -bind=10.211.55.12  -rejoin -join 10.211.55.11

./consul agent -server -bootstrap-expect 2 -data-dir /data/consul-data -node=n3 -bind=10.211.55.13  -rejoin -join 10.211.55.11
```

> `-server` 定义agent以server的模式运行
>
> `-bootstrap-expect`:在一个数据中心期待提供的server节点数目，当该值提供的时候，consul一直等到达到指定server数目的时候才会引导整个集群，该标记不能和`bootstrap`共用
>
> `-data-dir`:提供一个目录用来存储agent的状态，所有的agent允许都需要该目录，该目录必须是稳定存在的
>
> `-node`:节点在集群中的名称，在一个集群中必须是唯一的，默认是该节点的主机名
>
> `-bind`:该地址用来在集群内部的通讯，集群内的所有节点到该地址都必须是可达的默认是 0.0.0.0
>
> `-ui`: 启动web界面
>
> `-config-dir`: 配置文件的目录，里面所有以.json结尾的文件都会被加载
>
> `-rejoin`: 使consul忽略先前的离开，再次启动后仍旧尝试加入集群中
>
> `-client`:consul服务监听的地址，这个地址提供HTTP，DNS,RPC等服务，默认是127.0.0.1所以不对外提供服务，如果你要对提供服务改成0.0.0.0
>
> 

```
consul members  //查看集群成员

ctrl + c 终止会被检测为异常，标记为危险，会进行重连
consul leave 	//优雅离开
```

#### 注册服务

首先，为consul配置创建一个目录，consul会载入配置文件夹里的所有配置文件，在unix系统中通常类似/etc/consul.d（.d后缀意思是这个路径包含了一组配置文件）

```sh
mkdir /etc/consul.d
```

然后，编写服务定义配置文件，假设有一个web的服务运行在8080端口，可以给它设置一个标签，这样我们可以使用他作为额外的查询方式。

```json
{
  "service":{
    "name":"web",
    "tag":["master"],  #标记
    "address":"127.0.0.1",
    "port":8080, #服务的端口
    "checks":[
      {
        "http":"http://localhost:8080/health", #检查健康状态
        "interval":"10s"
      }
    ]
  }
}
```

![image-20220522213405634](https://gitee.com/ChetWei/img/raw/master/img/202205222134869.png)

> 我们只看数据中心1，可以看出consu的集群是由N个SERVER，加上M个CLIENT组成的。而不管是SERVER还是CLIENT，都是cosu的一个节点，所有的服务都可以注册到这些节点上，正是通过这些节点实现服务注册信息的共享。
>
> 除了这两个，还有一些小细节，一简单介绍。CLIENT CLIENT表示consul的client模式，就是客户端模式。是consul节点的一种模式，这种模式下，所有注册到当前节点的服务会被转发到SERVER【通过HTTP和DNS接口请求server】，本身是不持久化这些信息。SERVER SERVER表示consule的server模式，表明这个consul是个sevr，这种模式下，功能和CLIENT都一样，唯一不同的是，它会把所有的信息持久化的本地，这样遇到故障，信息是可以被保留的SERVER-LEADER中间那个SERVER下面有LEADER的字眼，表明这个SERVER是它们的老大，它和其它SERVER不一样的一点是，它需要负责同步注册的信息给其它的SERVER，同时也要负责各个节点的健康监测。

### Go-Micro

