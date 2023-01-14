package main

import (
	"fmt"
	"unsafe"
)

// golang的bool類型使用
func main() {
	var b = false
	fmt.Println("b=", b)

	//bool類型佔用一個字節
	fmt.Println("b佔用的空間=", unsafe.Sizeof(b))

	//報錯：b只能取true/false
	//b = 1
}
