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

	fmt.Println()

	//二維陣列使用方式1：先聲明/定義，再賦值
	var arr2 [2][3]int
	arr2[1][1] = 10
	fmt.Println("arr2=", arr2)

	fmt.Printf("arr2[0]的地址：%p \n", &arr2[0])
	fmt.Printf("arr2[1]的地址：%p \n", &arr2[1])

	fmt.Printf("arr2[0][0]的地址：%p \n", &arr2[0][0])
	fmt.Printf("arr2[1][0]的地址：%p \n", &arr2[1][0])

	fmt.Println()

	//二維陣列使用方式2：直接初始化
	arr3 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("arr3=", arr3)
}
