package main

import "fmt"

type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

// Deposit 存款
func (account *Account) Deposit(money float64, pwd string) {
	//看下輸入的密碼是否正確
	if pwd != account.Pwd {
		fmt.Println("你輸入的密碼不正確")
		return
	}

	//查看存款的金額是否正確
	if money <= 0 {
		fmt.Println("你輸入的金額不正確")
		return
	}

	account.Balance += money
	fmt.Println("存款成功")
}

// Withdraw 取款
func (account *Account) Withdraw(money float64, pwd string) {
	//看下輸入的密碼是否正確
	if pwd != account.Pwd {
		fmt.Println("你輸入的密碼不正確")
		return
	}

	//查看取款的金額是否正確
	if money <= 0 || money > account.Balance {
		fmt.Println("你輸入的金額不正確")
		return
	}

	account.Balance -= money
	fmt.Println("取款成功")
}

func (account *Account) Query(pwd string) {
	//看下輸入的密碼是否正確
	if pwd != account.Pwd {
		fmt.Println("你輸入的密碼不正確")
		return
	}

	fmt.Printf("你的帳號為%v 餘額為%v \n", account.AccountNo, account.Balance)
}

func main() {
	//測試
	account := &Account{
		AccountNo: "aa111111",
		Pwd:       "666",
		Balance:   100.0,
	}

	account.Query("666")

	account.Deposit(200.0, "666")
	account.Query("666")

	account.Withdraw(50.0, "666")
	account.Query("666")
}
