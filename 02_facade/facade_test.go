package _2_facade

/**
外观模式：为子系统中的一组接口提供一个一致的界面，Facade模式定义了一个高层接口，这个接口使得这一子系统更加容易维护
本质：封装交互，简化调用，符合："迪米特法则"（不和陌生人说话）
注意和中介者模式的区分
*/
import "testing"

func TestNoPattern(t *testing.T) {
	p := &Presentation{}
	p.Generate()

	b := &Business{}
	b.Generate()

	d := &DAO{}
	d.Generate()
}

func TestPattern(t *testing.T) {
	f := &Facade{}
	f.Generate()
}
