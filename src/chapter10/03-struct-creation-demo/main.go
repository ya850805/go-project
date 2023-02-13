package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	//方式1
	var p1 Person
	fmt.Println("p1=", p1)

	//方式2
	p2 := Person{"Jason", 26}
	//p2.Name = "Jason"
	//p2.Age = 26
	fmt.Println("p2=", p2)

	//方式3
	var p3 *Person = new(Person)
	//因為p3是一個指針，因此標準給字段賦值的方式是
	(*p3).Name = "Juan"
	//(*p3).Name = "Juan" 也等價於 p3.Name = "Juan"
	//原因：go的設計者為了工程師使用方便，底層會對p3.Name = "Juan"進行處理
	//會給p3加上取值運算(*p3).Name = "Juan"
	p3.Name = "Juan1"

	(*p3).Age = 27
	fmt.Println("*p3=", *p3)

	//方式4
	var p4 *Person = &Person{"M", 6}
	//因為p4是一個指針，因此標準給字段賦值的方式是
	(*p4).Name = "Marry"
	//簡化寫法
	p4.Name = "Marry1"

	(*p4).Age = 10
	fmt.Println("*p4=", *p4)
}
