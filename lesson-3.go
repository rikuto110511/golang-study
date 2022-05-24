package main

import (
	"fmt"
)

// 定数について

// const はアンタイム（型を宣言しない）
// 実行時に型を特定する

const Pi = 3.14

const (
	Username = "test_user"
	Password = "test_pass"
)

func main() {
	fmt.Println(Pi, Username, Password)

}
