package main

import "fmt"

func main() {
	var seconds float64
	var gender string

	fmt.Print("請輸入秒數：")
	fmt.Scanln(&seconds)

	fmt.Print("請輸入性別：")
	fmt.Scanln(&gender)

	if seconds <= 8 {
		if gender == "男" {
			fmt.Println("進入決賽男子組")
		} else {
			fmt.Println("進入決賽女子組")
		}
	} else {
		fmt.Println("淘汰了...")
	}
}
