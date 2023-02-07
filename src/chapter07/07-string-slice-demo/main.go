package main

import "fmt"

func main() {
	//`string`底層是一個`byte`陣列，因此`string`也可以做切片處理
	str := "hello@world"

	//使用切片獲取到"world"
	slice := str[6:]
	fmt.Println("slice=", slice)

	//**`string`是不可變的**，也就是說不能通過`str[0] = 'z'`的方式來修改字符串
	//str[0] = 'z' //編譯不通過，string是不可變的

	//如果需要修改字符串，可以先將`string`:point_right:`[]byte`或者`[]rune`:point_right:修改:point_right:重寫轉成`string`。
	//"hello@world" -> "zello@world"
	arr1 := []byte(str)
	arr1[0] = 'z'
	str = string(arr1)
	fmt.Println("str=", str)

	//我們轉成[]byte後，可以處理英文和數字，但是不能處理中文
	//原因是[]byte是字節處理，而一個中文字是3個字節，因此會出現亂碼
	//解決方法是將`string`轉成`[]rune`即可，因為`[]rune`是按字符處理，兼容中文字
	rune := []rune(str)
	rune[0] = '北'
	str = string(rune)
	fmt.Println("str=", str)
}
