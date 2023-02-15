package model

import "fmt"

type account struct {
	acctNo  string
	balance float64
	pwd     string
}

func NewAccount(accountNo string, balance float64, pwd string) *account {
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("帳號長度有誤")
		return nil
	}

	if balance <= 20 {
		fmt.Println("餘額有誤")
		return nil
	}

	if len(pwd) != 6 {
		fmt.Println("密碼長度有誤")
		return nil
	}

	return &account{accountNo, balance, pwd}
}

func (account *account) SetAccountNo(no string) {
	if len(no) >= 6 && len(no) <= 10 {
		account.acctNo = no
	} else {
		fmt.Println("帳號長度有誤")
	}
}

func (account *account) SetBalance(balance float64) {
	if balance > 20 {
		account.balance = balance
	} else {
		fmt.Println("餘額有誤")
	}
}

func (account *account) SetPwd(pwd string) {
	if len(pwd) == 6 {
		account.pwd = pwd
	} else {
		fmt.Println("密碼長度有誤")
	}
}

func (account *account) GetAccountNo() string {
	return account.acctNo
}

func (account *account) GetBalance() float64 {
	return account.balance
}

func (account *account) GetPwd() string {
	return account.pwd
}
