package learn

import (
	"fmt"
)

// class に相当する概念
// メンバ変数の定義のみ可能
type Item struct {
	Name  string
	Value int
	// private なメンバ変数は先頭を小文字にする
	password string
}

// メソッドの定義
// func (レシーバ : thisに相当) メソッド名(引数) 戻り値
func (item *Item) SetPerson(name string, value int) {
	item.Name = name
	item.Value = value
}

func Struct() {
	item := Item{}
	item.SetPerson("hoge", 100)

	fmt.Println(item.Name)
	fmt.Println(item.password) // setされていないためnil

	// 普通に宣言することも可能
	// ただし、すべて必須引数になるため、
	// 上記Setterを定義しておくと必要分のみ設定可能
	item2 := Item{"hoge2", 200, "password"}

	fmt.Println(item2.Name)

	item3 := Item{
		Name:     "hoge3",
		Value:    300,
		password: "password",
	}

	fmt.Println(item3.Name)
}
