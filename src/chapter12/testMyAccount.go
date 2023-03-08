package main

import "fmt"

func main() {
	//聲明一個變數，保存接收用戶收入的選項
	key := ""

	//聲明一個變數，控制是否退出for循環
	loop := true

	//帳戶餘額
	balance := 10000.0
	//每次收支的金額
	money := 0.0
	//每次收支的說明
	note := ""
	//定義一個變數，紀錄是否有收支的行為
	flag := false
	//收支的詳情，當有收支時，只需要拼接這個字符串即可
	details := "收支\t帳戶金額\t收支金額\t說明"

	//顯示主菜單
	for loop {
		fmt.Println("----------家庭收支記帳軟體----------")
		fmt.Println("          1. 收支明細")
		fmt.Println("          2. 登記收入")
		fmt.Println("          3. 登記支出")
		fmt.Println("          4. 退出軟體")
		fmt.Print("請選擇(1-4)：")
		fmt.Scanln(&key)

		switch key {
		case "1":
			fmt.Println("----------當前收支明細紀錄----------")
			if flag {
				fmt.Println(details)
			} else {
				fmt.Println("當前沒有任何收支明細")
			}
		case "2":
			fmt.Print("本次收入金額：")
			fmt.Scanln(&money)
			balance += money //修改餘額

			fmt.Print("本次收入說明：")
			fmt.Scanln(&note)

			//將收入情況拼接到details
			details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", balance, money, note)
			flag = true
		case "3":
			fmt.Print("本次支出金額：")
			fmt.Scanln(&money)
			//判斷餘額是否足夠
			if balance < money {
				fmt.Println("支出的金額不足")
				break //退出switch
			}

			balance -= money //修改餘額

			fmt.Print("本次支出說明：")
			fmt.Scanln(&note)

			//將收入情況拼接到details
			details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", balance, money, note)
			flag = true
		case "4":
			fmt.Println("你確定要退出嗎？y/n")
			choice := ""
			for {
				fmt.Scanln(&choice)
				if choice == "y" || choice == "n" {
					break
				} else {
					fmt.Println("輸入有誤，請重新輸入y/n")
				}
			}

			if choice == "y" {
				loop = false
			}

			fmt.Println("退出使用")
		default:
			fmt.Println("請輸入正確的選項")
		}

		fmt.Println()
		fmt.Println()
		fmt.Println()
	}
}
