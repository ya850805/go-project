package main

import "fmt"

func main() {
	//請求出一個陣列的最大值，並得到對應的下標
	array := [5]int{3, 5, 2, 7, 1}

	var max, index = array[0], 0
	for i := 1; i < len(array); i++ {
		if array[i] > max {
			max = array[i]
			index = i
		}
	}

	fmt.Printf("最大值為：%d, 下標為：%d", max, index)
}
