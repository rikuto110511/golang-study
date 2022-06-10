package lesson14

import "fmt"

// 可変長引数について

func foo(params ...int) {
	fmt.Println(len(params), params)

	for _, param := range params {
		fmt.Println(param)
	}
}

func lesson14() {
	foo()
	foo(10, 20)
	foo(10, 20, 30)
	foo(10, 20, 30, 40)

	s := []int{1, 2, 3}
	fmt.Println(s)

	foo(s...)
}
