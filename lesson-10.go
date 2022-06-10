package lesson10

import "fmt"

// mapについて

func lesson10() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["banana"] = 300
	fmt.Println(m)
	m["new"] = 500
	fmt.Println(m)

	fmt.Println(m["noting"])

	v, ok := m["apple"]
	fmt.Println(v, ok)

	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2)

	// 二つ受け取らなくてもOK
	v3 := m["banana"]
	fmt.Println(v3)

	m2 := make(map[string]int)
	m2["pc"] = 5000
	fmt.Println(m2)

	// 初期化していないためnilとなる
	// var m3 map[string]int
	// m3["pc"] = 5000
	// fmt.Println(m3)

	var s []int
	if s == nil {
		fmt.Println("Nil")
	}
}
