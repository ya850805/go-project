package main

import (
	"fmt"
	//為了使用utils.go文件的內容，我們需要先引入該包
	"go-project/src/chapter03/04-demo/model"
)

// 變數使用的注意事項
func main() {
	//**該區域**的**數據值**可以在**同一類型範圍內**不斷變化
	var i = 10
	i = 20
	i = 30
	//i = "name" //報錯，原因是不能改變數據類型
	fmt.Println("i=", i)

	//變數在同一個作用域(在一個函數或者在代碼塊)內不能重名
	var j = 99
	//var j = 100 //報錯，j已經在前方定義過
	fmt.Println("j=", j)

	//使用utils.go的HeroName，包名.標識符
	fmt.Println("hero name=", model.HeroName)
}
