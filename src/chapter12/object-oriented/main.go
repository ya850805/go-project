package main

import "fmt"

type FamilyAccount struct {
	key     string
	loop    bool
	balance float64
	money   float64
	note    string
	flag    bool
	details string
}

// NewFamilyAccount 編寫要給工廠模式的構造方法，返回一個*FamilyAccount實例
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "收支\t帳戶金額\t收支金額\t說明",
	}
}

//給FamilyAccount結構體綁定方法

// MainMenu 顯示主菜單
func (familyAccount *FamilyAccount) MainMenu() {
	for familyAccount.loop {
		fmt.Println("----------家庭收支記帳軟體----------")
		fmt.Println("          1. 收支明細")
		fmt.Println("          2. 登記收入")
		fmt.Println("          3. 登記支出")
		fmt.Println("          4. 退出軟體")
		fmt.Print("請選擇(1-4)：")
		fmt.Scanln(&familyAccount.key)

		switch familyAccount.key {
		case "1":
			familyAccount.ShowDetails()
		case "2":
			familyAccount.income()
		case "3":
			familyAccount.pay()
		case "4":
			familyAccount.exit()
		default:
			fmt.Println("請輸入正確的選項")
		}

		fmt.Println()
		fmt.Println()
		fmt.Println()
	}
}

// ShowDetails 顯示明細
func (familyAccount *FamilyAccount) ShowDetails() {
	fmt.Println("----------當前收支明細紀錄----------")
	if familyAccount.flag {
		fmt.Println(familyAccount.details)
	} else {
		fmt.Println("當前沒有任何收支明細")
	}
}

// income 登記收入
func (familyAccount *FamilyAccount) income() {
	fmt.Print("本次收入金額：")
	fmt.Scanln(&familyAccount.money)
	familyAccount.balance += familyAccount.money //修改餘額

	fmt.Print("本次收入說明：")
	fmt.Scanln(&familyAccount.note)

	//將收入情況拼接到details
	familyAccount.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", familyAccount.balance, familyAccount.money, familyAccount.note)
	familyAccount.flag = true
}

// pay 登記支出
func (familyAccount *FamilyAccount) pay() {
	fmt.Print("本次支出金額：")
	fmt.Scanln(&familyAccount.money)
	//判斷餘額是否足夠
	if familyAccount.balance < familyAccount.money {
		fmt.Println("支出的金額不足")
		return
	}

	familyAccount.balance -= familyAccount.money //修改餘額

	fmt.Print("本次支出說明：")
	fmt.Scanln(&familyAccount.note)

	//將收入情況拼接到details
	familyAccount.details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", familyAccount.balance, familyAccount.money, familyAccount.note)
	familyAccount.flag = true
}

// exit 退出
func (familyAccount *FamilyAccount) exit() {
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
		familyAccount.loop = false
	}

	fmt.Println("退出使用")
}

func main() {
	familyAccount := NewFamilyAccount()

	familyAccount.MainMenu()
}
