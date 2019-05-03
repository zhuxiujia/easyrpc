package main

import (
	"fmt"
	"github.com/zhuxiujia/easyrpc"
	"go-p2p/platform-common/src/com/platform/common/utils"
	"log"
	"net"
	"time"
	//"fmt"
	//"net/rpc/jsonrpc"
)

type TestVO struct {
	Name string `json:"name"`
}

//注意字段必须是导出
type Params struct {
	Width, Height int
}

type Rect struct {
	Area func(p Params, ret *TestVO) error
}

//func (r *Rect) Area(p *Params, ret *int) error {
//	*ret = p.Width + p.Height
//	println("do")
//	return nil
//}

func main() {
	go serverce()

	//连接远程rpc服务
	//这里使用Dial，http方式使用DialHTTP，其他代码都一样
	c, err := easyrpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second)

	var total = 10
	defer utils.CountMethodTps(float64(total), time.Now(), "ZmicroRpcClient")
	var ret = TestVO{}
	for i := 0; i < total; i++ {
		c.Call("Rect.Area", &Params{50, 100}, &ret)
		fmt.Println(ret.Name)
	}
}
func serverce() {
	rect := new(Rect)
	rect.Area = func(p Params, ret *TestVO) error {
		var t = TestVO{
			Name: "sdafasdf",
		}
		*ret = t
		println("do")
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
