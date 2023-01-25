package main

import "fmt"

// golang中的if-else
func main() {
	//編寫一個程序，可以輸入人的年齡，
	//如果這個人的年齡大於等於18歲，則輸出"你已經滿18歲，需要對自己的行為負責！"。
	//否則輸出"你的年齡不大這次放過你了"

	var age byte
	fmt.Print("請輸入年齡：")
	fmt.Scanln(&age)

	if age >= 18 {
		fmt.Println("你已經滿18歲，需要對自己的行為負責！")
	} else {
		fmt.Println("你的年齡不大這次放過你了")
	}
}
