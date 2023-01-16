package main

import "fmt"

// golang中標識符的使用
func main() {
	//Golang中嚴格區分大小寫
	//Golang中認為num和Num是不同的變數
	var num int = 10
	var Num int = 20
	fmt.Println("num=", num, "Num=", Num)

	//標識符不能包含空格
	//var ab c int = 10

	//_是空標識符，用於佔用
	//var _ int = 30
	//報錯：Cannot use '_' as a value
	//fmt.Println(_)

	//`int`不是保留關鍵字，但最好不要這樣使用
	var int int = 1
	fmt.Println("int=", int)
}
