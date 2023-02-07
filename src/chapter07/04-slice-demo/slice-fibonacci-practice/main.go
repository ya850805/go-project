package main

import "fmt"

//編寫一個函數fbn(n int)，要求完成：
//
//可以接收一個n int。
//能夠將費波那契數列放到切片中。

func fbn(n int) []uint {
	if n == 1 {
		return []uint{1}
	}
	if n == 2 {
		return []uint{1, 1}
	}

	slice := make([]uint, n)
	slice[0] = 1
	slice[1] = 1
	for i := 2; i < n; i++ {
		slice[i] = slice[i-1] + slice[i-2]
	}

	return slice
}

func main() {
	slice := fbn(10)
	fmt.Println("slice=", slice)
}
