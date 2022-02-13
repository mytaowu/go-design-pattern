package _1_factory

import (
	"testing"
)

func TestNewAPI(t *testing.T) {
	apiA := NewAPI(1) // Java中的静态方法就相当于Go中的普通方法（静态工厂）
	apiA.test1("哈哈，不要紧张，只是个测试而已！")

	apiB := NewAPI(2)
	apiB.test1("哈哈，不要紧张，只是个测试而已！")
}
