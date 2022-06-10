package lesson11

import "fmt"

// バイト型について

func lesson11() {
	b := []byte{72, 73}
	fmt.Println(b)

	fmt.Println(string(b))

	c := []byte("HI")
	fmt.Println(c)
	fmt.Println(string(c))
}
