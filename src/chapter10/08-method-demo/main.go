package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) test() {
	p.Name = "Juan"
	fmt.Println("test() name=", p.Name)
}

func main() {
	var p Person
	p.Name = "Jason"
	p.test()
	fmt.Println("p.Name=", p.Name)
}
