package main

import "fmt"

type A struct {
	Name string
	Age  int
}

type B struct {
	Name  string
	Score float64
}

type C struct {
	A
	B
}

type D struct {
	a A //有名結構體，組合關係
}

type Goods struct {
	Name  string
	Price float64
}

type Brand struct {
	Name    string
	Address string
}

type TV struct {
	Goods
	Brand
}

type TV2 struct {
	*Goods
	*Brand
}

func main() {
	var c C

	//報錯，因為匿名結構體A, B都有Name字段
	//c.Name = "tom"

	//需要指定匿名結構體來區分調用哪個匿名結構體的字段
	c.A.Name = "tom"

	var d D
	d.a.Name = "tom~"

	//嵌套匿名結構體後，也可以在創建結構體變數(實例)時，直接指定各個匿名結構體字段的值
	var tv = &TV{
		Goods: Goods{"電視機001", 5000.99},
		Brand: Brand{"品牌1", "地址..."},
	}
	fmt.Println("tv=", tv)

	tv2 := TV2{
		Goods: &Goods{"電視機002", 7000.99},
		Brand: &Brand{"品牌2", "地址..."},
	}
	fmt.Println(*tv2.Goods)
	fmt.Println(*tv2.Brand)
}
