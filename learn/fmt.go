package learn

import (
	"fmt"
	"reflect"
)

func Format() {
	num := 10

	str := "Hello, World!"
	str = "changeable string"

	// constant
	// not changeable
	const constant = 20

	// multiple declaration
	name, age := "Kim", 10

	fmt.Println(num)
	fmt.Println(str)
	fmt.Println(name, age)

	// type ------------------------------
	fmt.Println("learning : type ------------------------------")

	num1 := 10
	num2 := 20
	num3 := 1.14

	// fmrを使えば、%Tで型を表示できる
	fmt.Printf("%T\n", num1)

	// reflectを使えば、型を取れる
	fmt.Println(reflect.TypeOf(num1))

	fmt.Println(reflect.TypeOf(num1) == reflect.TypeOf(num2))
	fmt.Println(reflect.TypeOf(num1) == reflect.TypeOf(num3))

	// 新しい型を定義する
	// type <新しい型> <既存の型>
	type MyInt int
	// ()で囲むことで、複数の型を定義できる
	type (
		Name string
		Age  int
	)

	fmt.Println(reflect.TypeOf(MyInt(10)))
	fmt.Println(reflect.TypeOf(Name("Kim")))
	fmt.Println(Name("Kim"))
}
