

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