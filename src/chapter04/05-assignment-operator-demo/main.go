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

	var c int
	c = a + 3 //賦值的執行順序是從右向左
	fmt.Println("c=", c)

	//賦值運算符的**左邊只能是變數**，右邊可以是變數、表達式、常數
	//表達式：任何有值的都可以看做表達式
	var d int
	d = a
	d = 8 + 2*8
	d = 890
	d = test()
	fmt.Println("d=", d)
}

func test() int {
	return 90
}
