package main

import "fmt"

// Golang中字符類型使用
func main() {
	var c1 byte = 'a'
	var c2 byte = '0' //字符的0

	//當我們直接輸出byte值時，就是輸出了對應字符的ascii code
	fmt.Println("c1=", c1, "c2=", c2)

	//如果我們希望輸出對應的字符，需要使用格式化輸出
	fmt.Printf("c1=%c c2=%c \n", c1, c2)

	//中文字如果用byte存由於對應的ascii code過大，因此會報錯overflow
	var c3 = '北'
	fmt.Printf("c3=%c c3對應的ascii code=%d \n", c3, c3)
}
