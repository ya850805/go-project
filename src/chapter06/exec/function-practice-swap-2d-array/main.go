package main

import "fmt"

func swap(array [3][3]int) [3][3]int {
	for i := 0; i < 3; i++ {
		for j := i; j < 3; j++ {
			t := array[j][i]
			array[j][i] = array[i][j]
			array[i][j] = t
		}
	}
	return array
}

func main() {
	var array = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	array = swap(array)
	fmt.Println("array=", array)
}
