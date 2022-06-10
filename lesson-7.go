package lesson7

import (
	"fmt"
	"strconv"
)

// 型変換について

func lesson7() {
	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %f\n", xx, xx, xx)

	var y float64 = 1.2
	yy := int(y)
	fmt.Printf("%T %v %d\n", yy, yy, yy)

	var s string = "14"
	// z = int(s)
	// i, _ := strconv.Atoi(s)
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v\n", i, i)

	h := "Hello World"
	fmt.Println(string(h[0]))
}
