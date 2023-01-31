package main

import "fmt"

func main() {
	//*100以內的數求和，求出當和第一次大於20的當前數。*
	var sum = 0
	for i := 1; i <= 100; i++ {
		sum += i
		if sum > 20 {
			fmt.Println("和第一次大於20的當前數=", i)
			break
		}
	}

	//實現登陸驗證，有三次機會，如果用戶名為"Jason"、密碼為"666"提示登入成功，否則提示還有幾次機會。
	var chance = 3
	var username string
	var password string

	for {
		fmt.Print("請輸入用戶名：")
		fmt.Scanln(&username)
		fmt.Print("請輸入密碼：")
		fmt.Scanln(&password)

		chance--
		if username == "Jason" && password == "666" {
			fmt.Println("登入成功")
			break
		} else if chance == 0 {
			fmt.Println("超過3次登入失敗...")
			break
		} else {
			fmt.Printf("登入失敗，還有%d次登入機會 \n", chance)
		}
	}
}
