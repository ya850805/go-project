package main

import "fmt"

func main() {
	//有兩個變數`a`和`b`，要求將其值進行交換，但是**不允許使用中間變數**，最終打印結果
	var a int = 10
	var b int = 20

	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("a=%v, b=%v", a, b)
}
