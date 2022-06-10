package lesson2

import (
	"fmt"
)

// 変数宣言について

func lesson2() {
	// var i int = 1
	// var f64 float64 = 1.2
	// var s string = "test"
	// var t bool = true
	// var f bool = false
	// var t2, f2, bool = true, false

	// 関数の外でも宣言できる
	var (
		i   int     = 1
		f64 float64 = 1.2
		s   string  = "test"
		t   bool    = true
		f   bool    = false
	)
	fmt.Println(i, f64, s, t, f)

	// 関数内のみでしか宣言できない
	xi := 1
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)

	// 型表示
	fmt.Printf("%T\n", xf64)
	fmt.Printf("%T\n", xi)
}
