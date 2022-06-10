package lesson12

import "fmt"

// 関数について

/*
func add(x int, y int) {
	fmt.Println("add function")
	fmt.Println(x + y)
}
*/
func add(x int, y int) (int, int) {
	return x + y, x - y
}

/*
func cal(price, item int) (result int) {
	result = price * item
	return result
}
*/

func cal(price, item int) (result int) {
	result = price * item
	// 戻り値に変数名を宣言しておけばreturnのみでOK
	return
}

func convert(price int) float64 {
	return float64(price)
}

func lesson12() {
	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)

	r3 := cal(100, 2)
	fmt.Println(r3)

	// 関数も変数に代入できる
	f := func(x int) {
		fmt.Println("inner func", x)
	}

	f(1)

	// そのまんま実行する
	func(x int) {
		fmt.Println("inner func", x)
	}(2)
}
