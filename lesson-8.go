package lesson8

import (
	"fmt"
)

// 配列について

func lesson8() {
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	// 配列は追加できない
	/*
		var b [2]int = [2]int{100, 200}
		b = append(b, 300)
		fmt.Println(b)
	*/

	// スライスは追加できる
	var b []int = []int{100, 200}
	b = append(b, 300)
	fmt.Println(b)
}
