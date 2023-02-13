package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var p1 Person
	p1.Age = 10
	p1.Name = "小明"

	var p2 = &p1
	fmt.Println((*p2).Age)
	fmt.Println(p2.Age)

	p2.Name = "tom~"

	fmt.Printf("p2.Name=%v p1.Name=%v \n", p2.Name, p1.Name)
	fmt.Printf("p2.Name=%v p1.Name=%v \n", (*p2).Name, p1.Name)

	fmt.Printf("p1的地址是：%p \n", &p1)
	fmt.Printf("p2的地址是：%p \n", &p2)
	fmt.Printf("p2的值是：%p \n", p2)
}
