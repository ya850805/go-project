package main

import "fmt"

type Circle struct {
	radius float64
}

func (circle Circle) area() float64 {
	return 3.14 * circle.radius * circle.radius
}

func main() {
	//1. 聲明一個結構體`Circle`，字段為`radius`。
	//2. 聲明一個方法`area`和`Circle`綁定，可以返回面積。

	circle := Circle{4.0}
	res := circle.area()
	fmt.Println("面積為：", res)
}
