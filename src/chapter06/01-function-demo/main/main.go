package main

import (
	"fmt"
	"go-project/src/chapter06/01-function-demo/utils"
)

// golang中的函數
func main() {
	//輸入兩個數，再輸入一個運算符(+,-.\*,/)，得到結果。
	var num1, num2 float64
	var operator string

	fmt.Print("請輸入第一個數：")
	fmt.Scanln(&num1)
	fmt.Print("請輸入第二個數：")
	fmt.Scanln(&num2)
	fmt.Print("請輸入運算符(+,-,*,/)：")
	fmt.Scanln(&operator)

	//cal()函數可以使用多次，利於代碼重複使用
	res := utils.Cal(num1, num2, operator)

	fmt.Println("res=", res)
}
