package main

import "fmt"

type B interface {
	Test01()
}

type C interface {
	Test02()
}

type A interface {
	B
	C
	Test03()
}

// 如果需要實現A介面，就需要將B, C介面的方法都實現
type Stu struct {
}

func (s Stu) Test01() {

}
func (s Stu) Test02() {

}
func (s Stu) Test03() {

}

// 空介面
type T interface {
}

func main() {
	var stu Stu
	var a A = stu
	a.Test01()
	a.Test02()
	a.Test03()

	var t T = stu
	fmt.Println("t=", t)
	var t2 interface{} = stu
	var num1 float64 = 8.8
	t2 = num1
	fmt.Println("t2=", t2)
}
