package main

import (
	"context"
	"fmt"
	pd "gofiber05-gomicro/proto/hello"
	"google.golang.org/grpc"
)

func main() {

	//客户端连接服务器
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println("网络异常", err)
	}

	//网络延迟关闭
	defer conn.Close()
	//获得grpc句柄
	c := pd.NewHelloServiceClient(conn)

	//通过句柄远程调用函数
	re, err := c.SayHello(context.Background(), &pd.HelloReq{Name: "熊猫"})
	if err != nil {
		fmt.Println("远程调用SayHello失败")
	}
	fmt.Println("调用SayHello返回：", re.Msg)

	name, err := c.SayName(context.Background(), &pd.NameReq{Name: "托尼斯塔克"})
	if err != nil {
		fmt.Println("远程调用SayName失败")
	}

	fmt.Println("调用SayName返回：", name.Msg)

}
