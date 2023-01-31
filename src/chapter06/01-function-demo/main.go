package main

import "fmt"

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
	res := cal(num1, num2, operator)

	fmt.Println("res=", res)
}

// 將功能放到一個函數中，然後在需要使用時調用即可
func cal(num1 float64, num2 float64, operator string) float64 {
	var res float64

	switch operator {
	case "+":
		res = num1 + num2
	case "-":
		res = num1 - num2
	case "*":
		res = num1 * num2
	case "/":
		res = num1 / num2
	default:
		fmt.Println("操作錯誤")
	}

	return res
}
