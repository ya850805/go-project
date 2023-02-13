package main

import "fmt"

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp, rightDown Point
}

type Rect2 struct {
	leftUp, rightDown *Point
}

func main() {
	//結構體的所有字段在內存中是連續的。
	r1 := Rect{Point{1, 2}, Point{3, 4}}

	//r1有4個int，在內存中是連續分佈的
	fmt.Printf("r1.leftUp.x的地址：%p \n", &r1.leftUp.x)
	fmt.Printf("r1.leftUp.y的地址：%p \n", &r1.leftUp.y)
	fmt.Printf("r1.rightDown.x的地址：%p \n", &r1.rightDown.x)
	fmt.Printf("r1.rightDown.y的地址：%p \n", &r1.rightDown.y)

	fmt.Println()

	//r2有兩個*Point類型，這兩個*Point類型本身的地址也是連續的，但是他們指向的地址不一定是連續的
	r2 := Rect2{leftUp: &Point{10, 20}, rightDown: &Point{30, 40}}
	fmt.Printf("r2.leftUp本身地址：%p \n", &r2.leftUp)
	fmt.Printf("r2.rightDown本身地址：%p \n", &r2.rightDown)

	fmt.Println()

	//他們指向的地址不一定是連續的，這個要看系統運行時是如何分配的
	fmt.Printf("r2.leftUp指向的地址：%p \n", r2.leftUp)
	fmt.Printf("r2.rightDown指向的地址：%p \n", r2.rightDown)
}
