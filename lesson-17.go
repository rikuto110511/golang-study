package lesson17

import "fmt"

// for文について

func lesson17() {
	/*
		for i := 0; i < 10; i++ {
			if i == 3 {
				fmt.Println("continue")
				continue
			}

			if i > 5 {
				fmt.Println("break")
				break
			}
			fmt.Println(i)
		}
	*/

	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}

}
