package learn

import "fmt"

// init関数は、main関数よりも先に実行される
func init() {
	fmt.Println("init")
}

func InitFunction() {
	fmt.Println("initFunction")
}
