package main

import "fmt"

func main() {
	//使用for-range求一個陣列的和還有平均值
	array := [5]int{1, -1, 9, 90, 12}
	sum := 0.0

	for _, value := range array {
		sum += float64(value)
	}

	avg := sum / float64(len(array))

	fmt.Printf("和=%v, 平均值=%v", sum, avg)
}
