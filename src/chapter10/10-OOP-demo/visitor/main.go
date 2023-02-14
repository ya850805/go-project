package main

import "fmt"

type Visitor struct {
	Name string
	Age  int
}

func (visitor *Visitor) showPrice() {
	if visitor.Age > 18 {
		fmt.Println("門票20元")
	} else {
		fmt.Println("門票免費")
	}
}

func main() {
	var visitor Visitor
	for {
		fmt.Print("請輸入遊客的姓名：")
		fmt.Scanln(&visitor.Name)

		if visitor.Name == "n" {
			fmt.Println("退出程序...")
			break
		}
		fmt.Print("請輸入遊客的年明：")
		fmt.Scanln(&visitor.Age)

		visitor.showPrice()
	}
}
