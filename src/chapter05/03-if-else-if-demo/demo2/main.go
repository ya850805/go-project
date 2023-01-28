package main

import "fmt"

func main() {
	var b bool = true
	if b == false {
		fmt.Println("a")
	} else if b {
		fmt.Println("b")
	} else if !b {
		fmt.Println("c")
	} else {
		fmt.Println("d")
	}

	//如果第一個判斷改成`b = false`，會編譯錯誤，**因為`if`的條件表達式不能是賦值語句**。
}
