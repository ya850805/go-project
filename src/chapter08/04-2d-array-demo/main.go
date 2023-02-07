package main

import "fmt"

func main() {
	//二維陣列演示案例

	//定義/聲明一個二維陣列
	var arr [4][6]int

	//賦值
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}
