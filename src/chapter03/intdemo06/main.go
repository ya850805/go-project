package main

import "fmt"

// golang整數類型使用
func main() {
	var i int = 1
	fmt.Println("i=", i)

	//測試int8的範圍
	//報錯：The value of '-129' (type untyped int) cannot be represented by the type int8
	//var j int8 = -129

	//測試uint8的範圍
	//報錯：The value of '-1' (type untyped int) cannot be represented by the type uint8
	//var k uint8 = -1

	//int, uint, rune, byte的使用
	var a int = 8900
	fmt.Println("a=", a)

	//運行報錯：cannot use -1 (untyped int constant) as uint value in variable declaration (overflows)
	//var b uint = -1
	//fmt.Println("b=", b)

	//報錯：The value of '-1' (type untyped int) cannot be represented by the type byte
	//var c byte = -1

}
