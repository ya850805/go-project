package main

import "fmt"

// 一個被測試的函數
func addUpper(n int) int {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	return res
}

func main() {
	//傳統的測試方法，就是在main函數中使用看看結果是否正確
	res := addUpper(10)
	if res != 55 {
		fmt.Printf("addUpper錯誤，返回值=%v，期望值=%v", res, 55)
	} else {
		fmt.Printf("addUpper正確，返回值=%v，期望值=%v", res, 55)
	}
}
