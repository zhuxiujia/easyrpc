package main

import (
	"github.com/zhuxiujia/easyrpc"
	"github.com/zhuxiujia/easyrpc/example/test"
	"log"
	"net"
)

func main() {
	Serverce()
}

func Serverce() {
	rect := new(test.Rect)
	rect.Arg = func(arg string) (s string, e error) {
		println("do Arg:" + arg)
		return "do" + arg, nil
	}
	rect.NoArg = func() (s string, e error) {
		println("do NoArg")
		return "noarg", nil
	}
	rect.NoArgAndReturn = func() error {
		println("do NoArgAndReturn")
		return nil
	}
	rect.NoReturn = func(arg string) error {
		println("do NoReturn:arg=" + arg)
		return nil
	}

	//注册rpc服务
	easyrpc.Register(rect)
	var tcpUrl = "127.0.0.1:9999"

	l, e := net.Listen("tcp", tcpUrl)
	if e != nil {
		log.Fatalf("net rpc.Listen tcp :0: %v", e)
		panic(e)
	}
	for {
		conn, err3 := l.Accept()
		if err3 != nil {
			continue
		}
		//使用goroutine单独处理rpc连接请求
		go easyrpc.ServeConn(conn)
	}
}
