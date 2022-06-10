package lesson5

import (
	"fmt"
	"strings"
)

// 文字列型について

func lesson5() {
	fmt.Println("Hello world!!")
	fmt.Println("Hello " + "world!!")
	fmt.Println(string("Hello world!!"[0]))

	var s string = "Hello World"

	fmt.Println(strings.Replace(s, "H", "X", 1))
	fmt.Println(s)
	s = strings.Replace(s, "H", "X", 1)
	fmt.Println(s)
	fmt.Println(strings.Contains(s, "World"))

	fmt.Println(`Test
                     	Test

	Test`)

	fmt.Println("\"")
	fmt.Println(`"`)
}
