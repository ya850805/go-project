package main

import (
	"fmt"
)

func main() {

	var str string = "hello,world!你好"
	str2 := []rune(str)
	//字符串遍歷方式1 傳統方式
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c \n", str2[i])
	}

	//字符串遍歷方式2 for-range
	for index, val := range str {
		fmt.Printf("index=%d, val=%c \n", index, val)
	}
}
