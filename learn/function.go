package learn

import (
	"fmt"
)

func Function() {
	fmt.Println(add(42, 13))

	a, b := multiReturn("hello", "world")
	println(a, b)

	println(optionalArgs(1, 2, 3, 4, 5))
	println(optionalArgs(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

func add(x, y int) int {
	return x + y
}

// 複数の値を返却することもできる
func multiReturn(x, y string) (string, string) {
	return y, x
}

// 可変長引数を引数に定義できる
func optionalArgs(x, y int, z ...int) int {
	temp := x + y
	for _, v := range z {
		fmt.Println(v)
		temp += v
	}

	return temp
}
