package main

import (
	"fmt"
	"github.com/zhuxiujia/easyrpc"
	"log"
	"time"
	//"fmt"
	//"net/rpc/jsonrpc"
)

func main() {
	//连接远程rpc服务
	//这里使用Dial，http方式使用DialHTTP，其他代码都一样
	c, err := easyrpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second)

	do1(c)
	do2(c)
	do3(c)
	do4(c)
}

func do1(c *easyrpc.Client) {
	var result string
	var e = c.Call("Rect.NoArg", nil, &result)
	if e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Println("Rect.NoArg:" + result)
	}
}
func do2(c *easyrpc.Client) {
	var e = c.Call("Rect.NoReturn", "1", nil)
	if e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Println("Rect.NoReturn")
	}
}
func do3(c *easyrpc.Client) {
	var result string
	var e = c.Call("Rect.Arg", "1", &result)
	if e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Println("Rect.Arg:" + result)
	}
}
func do4(c *easyrpc.Client) {
	var e = c.Call("Rect.NoArgAndReturn", nil, nil)
	if e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Println("Rect.NoArgAndReturn")
	}
}
