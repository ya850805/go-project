package main

import "fmt"

func main() {
	var month byte
	var age byte
	var price float64 = 60

	fmt.Print("請輸入月份：")
	fmt.Scanln(&month)

	fmt.Print("請輸入年齡：")
	fmt.Scanln(&age)

	if month >= 4 && month <= 10 {
		if age >= 18 && age <= 60 {
			fmt.Println(price, "元")
		} else if age < 18 {
			fmt.Println(price/2, "元")
		} else {
			fmt.Println(price/3, "元")
		}
	} else {
		if age >= 18 && age <= 60 {
			fmt.Println("40元")
		} else {
			fmt.Println("20元")
		}
	}
}
