package main

import "fmt"

// BinarySearch 二分查找的函數
func BinarySearch(arr *[6]int, leftIndex int, rightIndex int, findVal int) {
	//判斷leftIndex是否大於rightIndex
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}

	middle := (leftIndex + rightIndex) / 2

	if (*arr)[middle] > findVal {
		//說明我們要查找的數在leftIndex~middle - 1
		BinarySearch(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		//說明我們要查找的數在middle + 1 - rightIndex
		BinarySearch(arr, middle+1, rightIndex, findVal)
	} else {
		//說明找到
		fmt.Printf("找到了，下標為%v \n", middle)
	}
}

func main() {
	arr := [6]int{1, 8, 10, 89, 1000, 1234}

	var num int
	fmt.Print("請輸入要查找的數字：")
	fmt.Scanln(&num)

	BinarySearch(&arr, 0, len(arr)-1, num)
}
