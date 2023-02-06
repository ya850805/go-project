package main

import "fmt"

// golang中的切片
func main() {
	var intArr [5]int = [...]int{11, 22, 33, 44, 55}

	//聲明/定義一個切片
	//1. slice就是切片名
	//2. intArr[1:3]表示slice引用到intArr這個陣列
	//3. 引用intArr陣列的起始下標為1，最後的下標為3(不包含)，因此是下標1~2的元素
	slice := intArr[1:3]
	fmt.Println("intArr的元素=", intArr)
	fmt.Println("slice的元素=", slice)
	fmt.Println("slice的元素個數=", len(slice))
	fmt.Println("slice的容量=", cap(slice)) //切片的容量是可以動態變化的

	fmt.Printf("intArr[1]的地址=%p \n", &intArr[1])
	fmt.Printf("slice[0]的地址=%p \n", &slice[0])
}
