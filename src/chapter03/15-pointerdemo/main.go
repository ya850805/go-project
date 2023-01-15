package main

import "fmt"

// golang中指針類型
func main() {
	//基本數據類型在內存佈局
	var i int = 10
	//i的地址是什麼？
	fmt.Println("i的地址是：", &i)

	//下面的 var ptr *int = &i
	//1. ptr是一個指針變數
	//2. ptr的類型是*int
	//3. ptr本身的值是&i
	var ptr *int = &i
	fmt.Println("ptr=", ptr)
	fmt.Println("ptr的地址是：", &ptr)
	fmt.Printf("ptr指向的值=%v", *ptr)
}
