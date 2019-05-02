

#EasyRPC 基于go标准库rpc框架修改而来
* 标准库默认使用func (* Type)Method(arg,*result) error 的模式,EasyRpc 则把方法移动到结构体里（方便动态代理和Aop以及各种扩展和定制）
* easyrpc同时支持 无参数，无返回值，或只有参数，只有返回值
```
type Service struct{
  Method func(arg ArgType,result *Result) error
}
``` 
```
type Service struct{
  Method func(arg ArgType) error
}
``` 
```
type Service struct{
  Method func(result *Result) error
}
``` 
```
type Service struct{
 Method func() error
}
``` 
# 使用方法

* 下载easyrpc
```
go get github.com/zhuxiujia/easyrpc
```
* 使用
```
//和go标准库的rpc以及jsonrpc使用方法完全一样,只需把rpc.* 和jsonrpc.* 改成 easyrpc.* 和 easy_jsonrpc.*
easyrpc.Client.Call()//client

easyrpc.Register(v) //rpc server
net.Listen("tcp", tcpUrl) //rpc server
```

# 搭配easyrpc_discovery使用
* https://github.com/zhuxiujia/easyrpc_discovery  //基于easyrpc定制微服务发现，支持动态代理，支持GoMybatis事务，AOP代理，事务嵌套，tag定义事务，自带负载均衡算法（随机，加权轮询，源地址哈希法）
![Image text](https://zhuxiujia.github.io/gomybatis.io/assets/easy_consul.png)
