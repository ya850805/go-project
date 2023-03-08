package main

import (
	"fmt"
	"go-project/src/chapter13/model"
	"go-project/src/chapter13/service"
)

type customerView struct {
	key  string //接收用戶輸入
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

// 得到用戶的輸入信息，構建新的客戶並完成添加
func (cv *customerView) add() {
	fmt.Println("----------添加客戶----------")
	fmt.Print("姓名：")
	name := ""
	fmt.Scanln(&name)

	fmt.Print("性別：")
	gender := ""
	fmt.Scanln(&gender)

	fmt.Print("年齡：")
	age := 0
	fmt.Scanln(&age)

	fmt.Print("電話：")
	phone := ""
	fmt.Scanln(&phone)

	fmt.Print("電郵：")
	email := ""
	fmt.Scanln(&email)

	//構建一個Customer
	//注意：id號沒有讓用戶輸入，id是唯一的，需要系統分配
	customer := model.NewCustomer2(name, gender, age, phone, email)

	if cv.customerService.Add(customer) {
		fmt.Println("----------添加完成----------")
	} else {
		fmt.Println("----------添加失敗----------")
	}
}

// 得到用戶輸入的id，刪除該id對應的客戶
func (cv *customerView) delete() {
	fmt.Println("----------刪除客戶----------")
	fmt.Print("請輸入要刪除的客戶編號(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return //放棄刪除操作
	}

	choice := ""
	for {
		fmt.Print("確認是否刪除(y/n)：")
		fmt.Scanln(&choice)
		if choice != "y" && choice != "Y" && choice != "n" && choice != "N" {
			fmt.Println("輸入錯誤，請重新數入")
		} else {
			break
		}
	}

	if choice == "y" || choice == "Y" {
		//調用customerService的Delete方法
		if cv.customerService.Delete(id) {
			fmt.Println("----------刪除完成----------")
		} else {
			fmt.Println("----------刪除失敗，輸入的id號不存在----------")
		}
	} else {
		fmt.Println("取消刪除操作")
	}
}

// 退出軟體
func (cv *customerView) exit() {
	for {
		fmt.Print("確認是否退出(y/n)：")
		fmt.Scanln(&cv.key)
		if cv.key == "y" || cv.key == "Y" || cv.key == "n" || cv.key == "N" {
			break
		}
		fmt.Println("輸入有誤請重新輸入")
	}

	if cv.key == "y" || cv.key == "Y" {
		cv.loop = false
		fmt.Println("退出客戶管理系統")
	}
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
			cv.add()
		case "2":
			fmt.Println("修改客戶")
		case "3":
			cv.delete()
		case "4":
			cv.list()
		case "5":
			cv.exit()
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
