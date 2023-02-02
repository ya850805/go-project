package main

import (
	"fmt"
)

// golang的可變參數
func sum(n1 int, nums ...int) int {
	for i := 0; i < len(nums); i++ {
		n1 += nums[i] //slice下標取值
	}

	return n1
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(6, 7, 7))
}
