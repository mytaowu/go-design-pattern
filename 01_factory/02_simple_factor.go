package _1_factory

import (
	"fmt"
)

type ImplA struct{}

func (i ImplA) test1(s string) {
	fmt.Println("Now in impA. the input s == ", s)
}

type ImplB struct{}

func (i ImplB) test1(s string) {
	fmt.Println("Now in impB. the input s == ", s)
}

// NewAPI 本质核心点是 "选择合适的实现类"
func NewAPI(i int) Api {
	if i == 1 { // 实际上会向客户端暴露出一些细节，也可以通过读取配置文件完成
		return ImplA{}
	}

	return ImplB{}
}
