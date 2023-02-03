package main

import "fmt"

// 函數外部聲明/定義的變數叫**全局變數**，作用域在**整個包都有效**，**如果其首字母為大寫，則作用域在整個程序有效**
var age int = 50
var Name string = "Jason"

func test() {
	//age和Name的作用域就只在test函數內部
	age := 10
	Name := "Tom"
	fmt.Println("age=", age)
	fmt.Println("Name=", Name)
}

func main() {
	fmt.Println("age=", age)
	fmt.Println("Name=", Name)

	//如果變數是在一個代碼塊(`for`/`if`)中，那麼這個變數的作用域就是帶該代碼塊
	for i := 0; i <= 10; i++ {
		fmt.Println("i=", i)
	}
}
