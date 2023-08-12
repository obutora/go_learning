package learn

import (
	"fmt"
)

func IfSwitch() {
	x := 10
	y := 20

	if y > x {
		fmt.Println("y is bigger than x")
	}

	if x > y {
		fmt.Println("x is bigger than y")
	} else {
		fmt.Println("x is smaller than ly")
	}

	text := "apple"

	// breakは不要
	switch text {
	case "apple":
		fmt.Println("text is apple")
	case "banana":
		fmt.Println("text is banana")
	}

	// 下の処理も実行する場合はfallthroughを使う
	switch text {
	case "apple":
		fmt.Println("text is apple")
		fallthrough
	case "banana":
		fmt.Println("text is banana")
	}
}
