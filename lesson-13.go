package lesson13

import "fmt"

// クロージャーについて

func incrementGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

// circleAreaの引数がXX
func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

func lesson13() {
	/*
			x := 0
			increment := func() int {
				x++
				return x
			}
				fmt.Println(increment())
		fmt.Println(increment())
		fmt.Println(increment())

	*/
	counter := incrementGenerator()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	c1 := circleArea(3.14)
	fmt.Println(c1(2))
	fmt.Println(c1(3))

	c2 := circleArea(3)
	fmt.Println(c2(2))
	fmt.Println(c2(3))
}
