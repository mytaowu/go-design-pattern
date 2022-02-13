package _1_factory

import "fmt"

type Api interface {
	test1(s string)
}

type Impl1 struct{}

func (i Impl1) test1(s string) {
	fmt.Println("Now in impl. the input s == ", s)
}
