package main

import (
	"fmt"
)

type A struct {
	Name string
	age  int
}

func (a *A) SayOk() {
	fmt.Println("A SayOk", a.Name)
}

func (a *A) hello() {
	fmt.Println("A hello", a.Name)
}

type B struct {
	A
	Name string
}

func (b *B) SayOk() {
	fmt.Println("B SayOk", b.Name)
}

func main() {
	//匿名結構體字段訪問可以簡化
	var b B
	b.Name = "tom"
	b.age = 19
	b.SayOk()
	b.hello()
}
