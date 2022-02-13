package _2_facade

type Facade struct{}

func (f *Facade) Generate() {
	p := &Presentation{}
	p.Generate()

	b := &Business{}
	b.Generate()

	d := &DAO{}
	d.Generate()
}
