package main

import "fmt"

func main() {

	var num int = 9
	fmt.Printf("num的地址=%v \n", &num)

	var ptr *int = &num
	*ptr = 10 //這裡修改時，會造成num值的變化
	fmt.Println("num=", num)

	////////////////////////////////
	var a int = 300
	var b int = 400
	var pointer *int = &a
	*pointer = 100
	pointer = &b
	*pointer = 200
	fmt.Printf("a=%d, b=%d, *pointer=%d, pointer=%v", a, b, *pointer, pointer)

}
