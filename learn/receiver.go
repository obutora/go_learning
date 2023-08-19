package learn

type Language struct {
	Name string
}

// (l *language)のことをreceiver引数という
// このように定義することで、Language構造体に対して、
// SetNameをメソッドとして追加できる
func (l *Language) SetName(name string) {
	l.Name = name
}

type MyInt int

func (f MyInt) IsZero() bool {
	return f == 0
}

func Printer() {
	l := Language{}
	l.SetName("Go")
	println(l.Name)

	l.SetName("Java")
	println(l.Name)

	mi := MyInt(2)
	println(mi.IsZero())

	mi = MyInt(0)
	println(mi.IsZero())
}

func Receiver() {
	Printer()
}
