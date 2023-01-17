package main

import "fmt"

// golang中&/*運算符的使用
func main() {
	a := 100
	fmt.Println("a的地址是：", &a)

	var ptr *int = &a
	fmt.Println("ptr的值(a的內存地址)是=", ptr)
	fmt.Println("ptr指向的值是：", *ptr)

	var n int
	var i int = 10
	var j int = 12
	//傳統的三元運算(go不支持三元運算)
	//n = i > j ? i : j
	if i > j {
		n = i
	} else {
		n = j
	}
	fmt.Println("n=", n)
}
