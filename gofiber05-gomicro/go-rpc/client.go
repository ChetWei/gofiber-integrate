package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	//建立网络连接
	cli, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("网络连接失败")
	}

	var pd int
	//传递对象.方法名称,参数，参数类型
	err = cli.Call("Panda.GetInfo", 10086, &pd)
	if err != nil {
		fmt.Println("调用失败!")
	}
	fmt.Println("远程调用方法得到的值", pd)
}
