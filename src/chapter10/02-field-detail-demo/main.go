package main

import "fmt"

// Person 如果結構體的字段是：指針、slice和`map`的默認值都是`nil`，即還沒有分配空間
// 如果需要使用這樣的字段需要先`make()`才能使用
type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int              //指針
	slice  []int             //切片
	map1   map[string]string //map
}

type Monster struct {
	Name string
	Age  int
}

func main() {
	//定義結構體變數
	var p1 Person
	fmt.Println("p1=", p1)

	fmt.Println("p1.ptr is nil?=>", p1.ptr == nil)
	fmt.Println("p1.slice is nil?=>", p1.slice == nil)
	fmt.Println("p1.map1 is nil?=>", p1.map1 == nil)

	//使用slice，需要先make
	p1.slice = make([]int, 10)
	p1.slice[0] = 100
	fmt.Println("p1=", p1)

	//使用map，需要先make
	p1.map1 = make(map[string]string)
	p1.map1["key1"] = "value1"
	fmt.Println("p1=", p1)

	fmt.Println()

	//**不同結構體變數的字段是獨立的，互不影響**，一個結構體變數字段更改，不影響另外一個。==**結構體是值類型**==。
	var monster1 Monster
	monster1.Name = "M1"
	monster1.Age = 500

	//結構體是值類型，所以是值拷貝
	monster2 := monster1
	monster2.Name = "M2"

	fmt.Println("monster1=", monster1)
	fmt.Println("monster2=", monster2)
}
