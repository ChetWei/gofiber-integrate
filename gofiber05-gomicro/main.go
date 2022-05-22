package main

import (
	"fmt"
	"gofiber05-gomicro/proto/testpackage"
	"google.golang.org/protobuf/proto"
)

func testsTest1() {
	//使用proto生成go文件
	test := &testpackage.Test{
		Name:   "panda",
		Weight: []int32{120, 135, 150, 180},
		Height: 180,
		Motto:  "世界和平!",
	}
	fmt.Println(test)

	//proto编码
	encodeData, err := proto.Marshal(test)
	if err != nil {
		fmt.Println("编码失败")
	}
	fmt.Println("编码结果:", encodeData)

	//proto解码
	newTest := &testpackage.Test{}
	err = proto.Unmarshal(encodeData, newTest)
	if err != nil {
		fmt.Println("解码失败")
	}
	fmt.Println("解码结果:", newTest)
}

func main() {

}
