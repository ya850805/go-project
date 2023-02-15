package main

import (
	"fmt"
	"go-project/src/chapter10/12-factory-demo/model"
)

func main() {
	//stu := model.Student{
	//	Name:  "tom",
	//	Score: 80.5,
	//}

	//fmt.Println(stu)

	//當student結構體首字母是小寫，我們可以通過工廠模式來解決
	var stu = model.NewStudent("tom~", 88.5)

	fmt.Println(*stu)
	fmt.Println("name=", stu.Name)
	//fmt.Println("score=", stu.Score)

	//使用GetScore取到score字段
	fmt.Println("score=", stu.GetScore())
}
