package utils

import "fmt"

// Cal 將功能放到一個函數中，然後在需要使用時調用即可
// 為了讓其他包的文件使用Cal函數，需要將函數首字母變成大寫，類似其他語言的public
func Cal(num1 float64, num2 float64, operator string) float64 {
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
