package main

import "fmt"

func main() {
	//根據用戶指定月份，打印該月份所屬的季節。3,4,5春季、6,7,8夏季、9,10,11秋季、12,1,2冬季。
	var month byte
	fmt.Print("請輸入月份：")
	fmt.Scanln(&month)

	switch month {
	case 3, 4, 5:
		fmt.Println("春季")
	case 6, 7, 8:
		fmt.Println("夏季")
	case 9, 10, 11:
		fmt.Println("秋季")
	case 12, 1, 2:
		fmt.Println("冬季")
	default:
		fmt.Println("輸入月份有誤...")
	}
}
