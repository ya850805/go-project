package main

import (
	"fmt"
	"go-project/src/chapter11/02-encapsulation-demo/practice/model"
)

func main() {
	account := model.NewAccount("123456", 500, "777777")

	if account != nil {
		fmt.Println("創建成功，account=", account)
	} else {
		fmt.Println("創建失敗")
	}

	account.SetAccountNo("1")
	fmt.Println("account number=", account.GetAccountNo())
}
