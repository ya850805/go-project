package main

import (
	"fmt"
	"go-project/src/chapter13/service"
)

type customerView struct {
	key  string //接收用戶選項輸入
	loop bool   //表示是否循環顯示主菜單

	customerService *service.CustomerService
}

// 顯示所有的客戶信息
func (cv *customerView) list() {
	//取得客戶列表
	customers := cv.customerService.List()
	//顯示
	fmt.Println("----------客戶列表----------")
	fmt.Println("編號\t姓名\t性別\t年齡\t電話\t郵箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("----------客戶列表完成----------")
}

// 顯示主菜單
func (cv *customerView) mainMenu() {
	for cv.loop {
		fmt.Println("----------客戶信息管理系統----------")
		fmt.Println("          1. 添加客戶")
		fmt.Println("          2. 修改客戶")
		fmt.Println("          3. 刪除客戶")
		fmt.Println("          4. 客戶列表")
		fmt.Println("          5. 退出")
		fmt.Print("請選擇(1-5)：")
		fmt.Scanln(&cv.key)

		switch cv.key {
		case "1":
			fmt.Println("添加客戶")
		case "2":
			fmt.Println("修改客戶")
		case "3":
			fmt.Println("刪除客戶")
		case "4":
			cv.list()
		case "5":
			cv.loop = false
			fmt.Println("退出客戶管理系統")
		default:
			fmt.Println("輸入錯誤，請重新輸入")
		}

		fmt.Println()
		fmt.Println()
		fmt.Println()
	}
}

func main() {
	customerView := customerView{
		key:             "",
		loop:            true,
		customerService: service.NewCustomerService(),
	}

	customerView.mainMenu()
}
