package main

import "fmt"

type AInterface interface {
	Say()
}

type BInterface interface {
	Hello()
}

type Stu struct {
	Name string
}

func (s Stu) Say() {
	fmt.Println("student say..")
}

type integer int

func (i integer) Say() {
	fmt.Println("integer say i =", i)
}

type Monster struct {
}

func (m Monster) Say() {
	fmt.Println("monster say()")
}

func (m Monster) Hello() {
	fmt.Println("monster hello()")
}

func main() {
	var stu Stu
	var a AInterface = stu
	a.Say()

	//只要是自定義數據類型，就可以實現介面，不僅僅是結構體類型
	var i integer = 10
	var b AInterface = i
	b.Say()

	//Monster實現了AInterface和BInterface
	var monster Monster
	var a2 AInterface = monster
	var b2 BInterface = monster
	a2.Say()
	b2.Hello()
}
