package main

import "fmt"

// BubbleSort 冒泡排序
func BubbleSort(arr *[5]int) {
	fmt.Println("排序前=", *arr)

	for i := 1; i <= len(arr)-1; i++ {
		for j := 0; j < len(arr)-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				//交換
				temp := (*arr)[j+1]
				(*arr)[j+1] = (*arr)[j]
				(*arr)[j] = temp
			}
		}
	}

	fmt.Println("排序後=", *arr)
}

func main() {
	arr := [5]int{24, 69, 80, 57, 13}

	//將陣列傳遞給一個函數，完成排序
	BubbleSort(&arr)

	fmt.Println("main arr=", arr)
}
