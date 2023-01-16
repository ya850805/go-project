package main

import "fmt"

// golang中邏輯運算符
func main() {

	//邏輯"與"的使用
	var age int = 40
	if age > 30 && age < 50 {
		fmt.Println("ok1")
	}
	if age > 30 && age < 40 {
		fmt.Println("ok2")
	}

	//邏輯"或"的使用
	if age > 30 || age < 50 {
		fmt.Println("ok3")
	}
	if age > 30 || age < 40 {
		fmt.Println("ok4")
	}

	//邏輯"非"的使用
	if age > 30 {
		fmt.Println("ok5")
	}
	if !(age > 30) {
		fmt.Println("ok6")
	}

	//短路與：**如果第一個條件為`false`，則第二個條件不會判斷，最終結果為`false`**
	var i int = 10
	if i < 9 && test() {
		fmt.Println("ok7")
	}

	//短路或：**如果第一個條件為`true`，則第二個條件不會判斷，最終結果為`true`**
	if i > 9 || test() {
		fmt.Println("ok8")
	}
}

// 測試函數
func test() bool {
	fmt.Println("test...")
	return true
}
