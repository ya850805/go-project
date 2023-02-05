package main

import "fmt"

// 觀察陣列&元素的地址
func main() {
	var intArr [2]int
	//當定義完陣列後，其實陣列各個元素有默認值
	fmt.Println("intArr=", intArr)
	intArr[0] = 10
	intArr[1] = 20
	fmt.Println("intArr=", intArr)

	fmt.Printf("intArr的地址=%p \n", &intArr)
	fmt.Printf("intArr第一個元素intArr[0]的地址=%p \n", &intArr[0])
	fmt.Printf("intArr第二個元素intArr[1]的地址=%p \n", &intArr[1])
}
