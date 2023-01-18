package main

import "fmt"

func main() {
	//要求可以從控制台獲取用戶信息(姓名、年齡、薪水、是否通過考試)
	//使用`fmt.Scanln()`
	var name string
	var age byte
	var sal float32
	var isPass bool

	fmt.Print("請輸入姓名：")
	//當程序執行到`fmt.Scanln()`會停止，等待用戶輸入，並回車
	fmt.Scanln(&name)

	fmt.Print("請輸入年齡：")
	fmt.Scanln(&age)

	fmt.Print("請輸入薪水：")
	fmt.Scanln(&sal)

	fmt.Print("請輸入是否通過考試：")
	fmt.Scanln(&isPass)

	fmt.Printf("姓名：%v, 年齡：%v, 薪水：%v, 是否通過考試：%v", name, age, sal, isPass)
}
