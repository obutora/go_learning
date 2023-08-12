package learn

import (
	"fmt"
)

// 大文字で始めた関数のみ公開される
func Array() {
	fmt.Println("learning : array ------------------------------")

	// array - 固定長
	a1 := [2]int{}
	a1[0] = 100
	a1[1] = 200

	fmt.Println(a1)

	// 宣言時の初期化も可能
	a2 := [2]int{100, 200}
	fmt.Println(a2)

	// 初期化時に要素数が固定される場合は...で省略可能
	a3 := [...]int{100, 200, 300}
	fmt.Println(a3)

	// slice - 可変長
	s1 := []int{100, 200, 300}
	fmt.Println(s1)

	// appendで要素を追加できる
	// appendは元の配列を変更するのではなく、新しい配列を作成して返す
	s1 = append(s1, 400)
	fmt.Println(s1)

	// sliceの要素を取得する
	fmt.Println(s1[0])

	// sliceの要素を更新する
	s1[0] = 1000
	fmt.Println(s1)

	// loop
	for i, v := range s1 {
		fmt.Println(i, v)
	}

	// map
	m1 := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m1)
	fmt.Println(m1["apple"])

	// mapに要素を追加する
	m1["new"] = 300
	fmt.Println(m1)

	// mapの要素を削除する
	delete(m1, "banana")
	fmt.Println(m1)

	// mapの要素数を取得する
	fmt.Println(len(m1))

	// loop処理
	for key, value := range m1 {
		fmt.Println(key, value)
	}

	// mapの要素が存在するか確認する
	// 要素が存在しない場合、エラーになる
	_, ok := m1["banana"]
	if ok {
		fmt.Println("exist")
	} else {
		fmt.Println("not exist")
	}
}
