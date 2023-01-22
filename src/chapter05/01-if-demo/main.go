package main

import "fmt"

// golang中的if
func main() {
	//編寫一個程序，可以輸入人的年齡，
	//如果這個人的年齡大於等於18歲，則輸出"你已經滿18歲，需要對自己的行為負責！"

	var age byte

	fmt.Print("請輸入年齡：")
	fmt.Scanln(&age)

	if age >= 18 {
		fmt.Println("你已經滿18歲，需要對自己的行為負責！")
	}

	//golang支持在if中直接定義一個變數
	if age1 := 21; age1 > 20 {
		fmt.Println("age1大於20")
	}
}
