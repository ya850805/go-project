package main

import "fmt"

func main() {
	//使用常規`for`循環遍歷切片
	var arr = [...]int{10, 20, 30, 40, 50}
	slice := arr[1:4]
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%d]=%v  \t", i, slice[i])
	}

	fmt.Println()

	//使用`for-range`
	for i, v := range slice {
		fmt.Printf("index=%d, value=%v \t", i, v)
	}
}
