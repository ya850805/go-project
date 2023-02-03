package main

import "fmt"

// golang的匿名函數

var (
	// Func1 func1就是一個全局匿名函數
	Func1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	//在定義匿名函數時就直接調用
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)

	fmt.Println("res1=", res1)

	//在匿名函數賦給一個變數(函數變數)，再通過該變數來調用匿名函數。
	//則a的數據類型就是函數類型
	a := func(n1 int, n2 int) int {
		return n1 - n2
	}

	res2 := a(10, 30)
	fmt.Println("res2=", res2)

	//全局匿名函數的使用
	res3 := Func1(4, 9)
	fmt.Println("res3=", res3)
}
