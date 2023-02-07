package main

import "fmt"

// 演示二維陣列的遍歷
func main() {
	arr := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	//雙層`for`循環
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%v \t", arr[i][j])
		}
		fmt.Println()
	}

	fmt.Println()

	//for-range方式完成遍歷
	for i, array := range arr {
		for j, num := range array {
			fmt.Printf("arr[%d][%d]=%v \t", i, j, num)
		}
		fmt.Println()
	}
}
