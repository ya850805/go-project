package main

import "fmt"

// 函數中的變數是局部的，函數外不生效
func test() {
	//n1是test函數的局部變數，只能在test函數中使用
	var n1 = 10
	fmt.Println("n1=", n1)
}

func test2(n1 *int) {
	*n1 += 10
	fmt.Println("test2 n1=", *n1)
}

func main() {
	//這裡不能使用n1，因為n1是test函數的局部變數
	//fmt.Println("n1=", n1)

	num := 20
	test2(&num)
	fmt.Println("main num=", num)
}
