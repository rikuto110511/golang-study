package lesson6

import (
	"fmt"
)

// 論理値型について

func lesson6() {
	t, f := true, false
	fmt.Printf("%T %v %t\n", t, 1, t)
	fmt.Printf("%T %v %t\n", f, 0, f)

	// かつ
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && false)

	// または
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || false)

	// 反転
	fmt.Println(!true)
	fmt.Println(!false)
}
