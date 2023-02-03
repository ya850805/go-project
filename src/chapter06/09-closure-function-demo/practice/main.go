package main

import (
	"fmt"
	"strings"
)

// 1. 編寫一個程序`makeSuffix(suffix string)`，可以接收一個文件後綴名(例如.jpg)，並返回一個閉包。
// 2. 調用閉包，可以傳入一個文件名，如果該文件名沒有指定的後綴(例如.jpg)，則返回文件名.jpg，如果已經有.jpg後綴，則返回原文件名。
// 3. 要求使用閉包的方式完成。
// 4. 可以使用`strings.HasSuffix()`來判斷後綴。
func makeSuffix(suffix string) func(string) string {
	return func(fileName string) string {
		if strings.HasSuffix(fileName, suffix) {
			return fileName
		} else {
			return fileName + suffix
		}
	}
}

func main() {
	function := makeSuffix(".jpg")

	fileName1 := "image.jpg"
	res1 := function(fileName1)
	fmt.Println("res1=", res1)

	fileName2 := "test"
	res2 := function(fileName2)
	fmt.Println("res2=", res2)
}
