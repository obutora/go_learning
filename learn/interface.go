package learn

import (
	"fmt"
)

// goのインターフェースは、メソッドの集まり
type Animal interface {
	Sing()
}

// インターフェースを実装しているかどうかをextendsなどで指定する必要はない
// interfaceで定義されたメソッドを実装していれば、そのインターフェースを実装しているとみなされる
type Cat struct{}

func (c Cat) Sing() {
	fmt.Println("Nyaa~")
}

type Dog struct{}

func (d Dog) Sing() {
	fmt.Println("Wan!")
}

func InterFace() {
	var animal Animal
	animal = Cat{}
	animal.Sing()
	animal = Dog{}
	animal.Sing()

	// インターフェースの型を調べる

	// animal の型は↑でDogである
	fmt.Printf("%T\n", animal)

	// Dog が Animal インターフェースを実装しているかどうかを調べる
	_, ok := animal.(Animal)
	fmt.Println(ok)
}
