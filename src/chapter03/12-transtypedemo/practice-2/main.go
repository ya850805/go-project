package main

import "fmt"

func main() {
	//練習題2
	var n1 int32 = 12
	//var n2 int8
	var n3 int8

	//編譯通過，但由於會溢出，因此結果不是127 + 12
	n3 = int8(n1) + 127
	fmt.Println("n3=", n3)

	//編譯失敗，因為128本身就超出n2 int8的範圍
	//n2 = int8(n1) + 128
}
