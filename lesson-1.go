package main

import (
	"fmt"
	"os/user"
	"time"
)

// importについて

/*
// initはmain関数で呼び出さなくても実行時に一番初めに実行される
func init() {
	fmt.Println("Init!")
}
*/

func bazz() {
	fmt.Println("Bazz")
}

func main() {
	// bazz()
	fmt.Println("hello world!", time.Now())
	fmt.Println(user.Current())
}
