package main

import "fmt"

// Golang沒有`while`和`do...while`語法，如果我們需要使用類似其他語言(ex.java的`while`/`do...while`)，可以通過`for`循環來實現其使用效果。
func main() {
	//使用類似`while`的語法輸出10句`hello world`
	var i = 1
	for {
		if i > 10 {
			break
		}

		fmt.Println("hello world", i)
		i++
	}

	//使用類似`do...while`的語法輸出10句`hello world!!!`
	var j = 1
	for {
		fmt.Println("hello world!!!", j)
		j++
		if j > 10 {
			break
		}
	}
}
