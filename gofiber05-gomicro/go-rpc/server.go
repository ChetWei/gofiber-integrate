package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/rpc"
)

func pandatest(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world")
}

//go rpc go语言自带的远程调用库

/*
- 方法是导出的
- 方法有两个参数，都是导出类型或内建类型
- 方法的第二个参数是指针
- 方法只有一个error接口类型的返回值
*/

type Panda int

//对端发送过来的内容，返回给对端的内容
func (this *Panda) GetInfo(argType int, replyType *int) error {
	fmt.Println("打印对端口号发送过来的内容:", argType)
	//修改内容值
	*replyType = argType + 12306

	return nil
}

func main() {
	//请求方法
	http.HandleFunc("/panda", pandatest)

	//将类实例化为对象
	pd := new(Panda)
	//服务端注册一个对象，使它作为服务暴露
	rpc.Register(pd)

	rpc.HandleHTTP()

	//开启监听
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("网络错误")
	}
	http.Serve(ln, nil)

}
