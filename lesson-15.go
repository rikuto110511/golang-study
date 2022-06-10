package lesson15

import "fmt"

// 演習
/*
Q1. 以下の1.11をint型に変換して出力してください。

f := 1.11

Q2. コードを書かずに以下の出力結果を答えてください。

s := []int{1, 2, 5, 6, 2, 3, 1}
fmt.Println(s[2:4])
Q3. 以下のコードを実行した時に

fmt.Printf("%T %v", m, m)

以下のような出力結果となるmを作成してください。

map[string]int map[Mike:20 Nancy:24 Messi:30]
*/

func lesson15() {
	// Q1
	f := 1.11
	fmt.Println(int(f))

	// Q2
	// 5,6
	s := []int{1, 2, 5, 6, 2, 3, 1}
	fmt.Println(s[2:4])

	// Q3
	m := make(map[string]int)
	m["Mike"] = 20
	m["Nancy"] = 24
	m["Messi"] = 30
	fmt.Printf("%T %v", m, m)

	m2 := map[string]int{
		"Mike":   20,
		"Naincy": 24,
		"Messi":  30,
	}
	fmt.Printf("%T %v", m2, m2)
}
