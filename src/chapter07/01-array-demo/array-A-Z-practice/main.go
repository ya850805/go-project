package main

import "fmt"

func main() {
	//創建一個`byte`類型的26個元素的陣列，分別放置'A'-'Z'。使用`for`循環訪問所有元素並打印出來
	var array [26]byte
	for i := 0; i < len(array); i++ {
		array[i] = 'A' + byte(i)
	}

	//print
	for _, ch := range array {
		fmt.Printf("%c ", ch)
	}
}
