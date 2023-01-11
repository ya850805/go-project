package main

import (
	"fmt"
	"unsafe"
)

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

	var n1 = 100 //n1是什麼類型
	//fmt.Printf() 可以用於做格式化輸出
	fmt.Printf("n1的類型是%T \n", n1)

	//如何在程序中查看某個變數的佔用字節大小和數據類型
	var n2 int64 = 10
	//unsafe.Sizeof() 是unsafe包裡面提供的函數，可以返回變數佔用的字節數
	fmt.Printf("n2的類型是%T，n2佔用的字節數是%d", n2, unsafe.Sizeof(n2))
}
