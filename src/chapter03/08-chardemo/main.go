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

	//**字符的本質是一個整數**，直接輸出時，是該字符對應的UTF-8編碼的碼值
	var c4 byte = '\n'
	fmt.Println("c4=", c4)

	//格式化%c輸出某個數，會輸出該數字對應的unicode字符
	var c5 = 21271
	fmt.Printf("c5=%c \n", c5)

	//字符型是可以進行運算的，相當於一個整數，運算時是按照碼值運算
	var n1 = 10 + 'a'
	fmt.Printf("n1=%c", n1)
}
