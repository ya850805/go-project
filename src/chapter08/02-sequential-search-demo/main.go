package main

import "fmt"

func main() {
	//有一個列表：A, B, C, D
	//從鍵盤任意輸入一個大寫字母，判斷列表是否包含此字母(順序查找)
	arr := [4]string{"A", "B", "C", "D"}

	var letter string
	fmt.Print("請輸入一個大寫字母：")
	fmt.Scanln(&letter)

	//順序查找(第一種)
	for i := 0; i < len(arr); i++ {
		if letter == arr[i] {
			fmt.Printf("找到%v, 下標為%d \n", letter, i)
			break
		} else if i == len(arr)-1 {
			fmt.Printf("沒有找到%v... \n", letter)
		}
	}

	//順序查找(第二種)(較推薦)
	index := -1
	for i := 0; i < len(arr); i++ {
		if letter == arr[i] {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Printf("沒有找到%v... \n", letter)
	} else {
		fmt.Printf("找到%v, 下標為%d \n", letter, index)
	}
}
