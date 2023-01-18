package main

import "fmt"

func main() {
	//要求可以從控制台獲取用戶信息(姓名、年齡、薪水、是否通過考試)
	//使用`fmt.Scanf()`
	var name string
	var age byte
	var sal float32
	var isPass bool

	fmt.Println("請輸入姓名、年齡、薪水、是否通過考試，並使用空格隔開")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf("姓名：%v, 年齡：%v, 薪水：%v, 是否通過考試：%v", name, age, sal, isPass)
}
