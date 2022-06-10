package lesson16

import "fmt"

// if文について

func by2(num int) string {

	if num%2 == 0 {
		return "ok"
	} else {
		return "no"
	}
}

func lesson16() {
	/*
		num := 9
		if num%2 == 0 {
			fmt.Println("by 2")
		} else if num%3 == 0 {
			fmt.Println("by 3")
		} else {
			fmt.Println("else")
		}
	*/

	result := by2(2)
	if result == "ok" {
		fmt.Println("Great")
	}
	fmt.Println(result)

	if result2 := by2(10); result2 == "ok" {
		fmt.Println("Great2")
	}

	x, y := 11, 12
	if x == 10 && y == 10 {
		fmt.Println("&&")
	}

	if x == 10 || y == 10 {
		fmt.Println("||")
	}
}
