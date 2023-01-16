package main

import "fmt"

// golang中的賦值運算符
func main() {
	var i int
	//基本賦值
	i = 10
	fmt.Println("i=", i)

	//有兩個變數a b，要求將其進行交換，最終打印結果
	a := 9
	b := 2
	fmt.Printf("交換前的情況是a=%v, b=%v \n", a, b)

	//定義一個臨時變數
	t := a
	a = b
	b = t
	fmt.Printf("交換後的情況是a=%v, b=%v \n", a, b)

	//複合賦值的操作
	a += 17 // 等價a = a + 17
	fmt.Println("a=", a)
}
