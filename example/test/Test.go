package test

type TestVO struct {
	Name string `json:"name"`
}

//注意字段必须是导出
type Params struct {
	Width, Height int
}

type Rect struct {
	NoArg          func() (string, error)           //无参数写法
	Arg            func(arg string) (string, error) //有参数有返回值
	NoReturn       func(arg string) error           //无返回值
	NoArgAndReturn func() error                     //无参数无返回值
}
