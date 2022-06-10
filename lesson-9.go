package lesson9

import "fmt"

// スライスについて

func lesson9() {
	// n := []int{1, 2, 3, 4, 5, 6}
	// fmt.Println(n)
	// fmt.Println(n[2])
	// fmt.Println(n[2:4])
	// fmt.Println(n[:2])
	// fmt.Println(n[2:])
	// fmt.Println(n[:])

	// n[2] = 100
	// fmt.Println(n)

	// var board = [][]int{
	// 	[]int{0, 1, 2},
	// 	[]int{3, 4, 5},
	// 	[]int{6, 7, 8},
	// }
	// fmt.Println(board)

	// n = append(n, 100)
	// fmt.Println(n)

	// n = append(n, 200, 300, 400)
	// fmt.Println(n)

	// 第一引数：配列、第二引数：初期サイズ（０で初期化）、第三引数：長さ
	n := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	n = append(n, 0, 0)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	n = append(n, 0, 0, 0, 0, 0)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)

	a := make([]int, 3)
	fmt.Printf("len=%d cap=%d value=%v\n", len(a), cap(a), a)

	// 0で初期化
	b := make([]int, 0)
	// nilなのでメモリ領域に確保しない
	var c []int
	fmt.Printf("len=%d cap=%d value=%v\n", len(b), cap(b), b)
	fmt.Printf("len=%d cap=%d value=%v\n", len(c), cap(c), c)

	// c = make([]int, 5)
	c = make([]int, 0, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(c), cap(c), c)

	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)
}
