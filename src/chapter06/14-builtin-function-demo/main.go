package main

import "fmt"

// golang的內置函數
func main() {
	num1 := 100
	fmt.Printf("num1的類型：%T, num1的值：%v, num1的地址：%v \n", num1, num1, &num1)

	num2 := new(int)
	fmt.Printf("num2的類型：%T, num2的值：%v, num2的地址：%v, num2指針對應的值是：%v \n", num2, num2, &num2, *num2)
}
