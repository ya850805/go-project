package main

import "fmt"

func modifyMap(map1 map[int]int) {
	map1[10] = 900
}

// Stu 定義一個學生結構體
type Stu struct {
	Name    string
	Age     int
	Address string
}

func main() {
	//`map`是引用類型，遵守引用類型傳遞的機制，在一個函數接收`map`修改後，會直接修改原來的`map`
	map1 := make(map[int]int)
	map1[1] = 90
	map1[2] = 88
	map1[10] = 1
	map1[20] = 2
	fmt.Println("map1=", map1)
	//map1會做修改，說明map是引用類型
	modifyMap(map1)
	fmt.Println("map1=", map1)

	fmt.Println()

	//`map`的value也經常使用`struct`類型，更適合管理複雜的數據
	students := make(map[string]Stu)
	//創建兩個學生
	stu1 := Stu{Name: "Jason", Age: 26, Address: "嘉義"}
	stu2 := Stu{Name: "Juan", Age: 27, Address: "台中"}
	students["001"] = stu1
	students["002"] = stu2
	fmt.Println("students=", students)

	//遍歷各個學生的信息
	for k, v := range students {
		fmt.Printf("學生的編號是：%v, 姓名是：%v, 年齡是：%v, 地址是：%v \n", k, v.Name, v.Age, v.Address)
	}
}
