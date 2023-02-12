package main

import "fmt"

// Cat 定義一個Cat結構體，將Cat的各個字段/屬性信息，放入到Cat結構體進行管理
type Cat struct {
	Name  string
	Age   int
	Color string
}

func main() {
	//小明養了兩隻貓：一隻名字叫小白，今年3歲，白色。還有一隻叫小橘，今年10歲，橘色。
	//請編寫一個程序，當用戶輸入小貓的名字時，就顯示該貓的名字、年齡、顏色，如果用戶輸入的名稱錯誤，就顯示小明每有這隻貓。

	//單獨的定義變數解決
	//var cat1Name = "小白"
	//var cat1Age = 3
	//var cat1Color = "白色"
	//
	//var cat2Name = "小橘"
	//var cat2Age = 10
	//var cat2Color = "橘色"

	//使用陣列解決
	//var catNames = [...]string{"小白", "小橘"}
	//var catAges = [...]int{3, 10}
	//var catColors = [...]string{"白色", "橘色"}

	//使用struct結構體來完成案例

	//創建一個Cat的變數
	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	fmt.Printf("cat1的地址是：%p \n", &cat1)

	fmt.Println("cat1=", cat1)
	fmt.Println("cat1 name=", cat1.Name)
	fmt.Println("cat1 age=", cat1.Age)
	fmt.Println("cat1 color=", cat1.Color)
}
