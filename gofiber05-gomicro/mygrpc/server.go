package main

import (
	"context"
	"fmt"
	pd "gofiber05-gomicro/proto/hello" //导入proto生成的go文件
	"google.golang.org/grpc"
	"net"
)

//google 的 mygrpc

type server struct {
}

//grpc
//将pb文件的方法进行实现
func (this *server) SayHello(ctx context.Context, in *pd.HelloReq) (out *pd.HelloResp, err error) {
	fmt.Println("SayHello handler")
	return &pd.HelloResp{Msg: "hello" + in.Name}, nil
}

func (c *server) SayName(ctx context.Context, in *pd.NameReq) (out *pd.NameResp, err error) {
	fmt.Println("SayName handler")
	return &pd.NameResp{Msg: in.Name + "早上好"}, nil
}

func main() {

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("网络错误", err)
	}

	//创建grpc的服务
	srv := grpc.NewServer()
	//注册服务
	pd.RegisterHelloServiceServer(srv, &server{})

	err = srv.Serve(ln)
	if err != nil {
		fmt.Println("网络错误", err)
	}
}
