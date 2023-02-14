package main

import "fmt"

type Circle struct {
	radius float64
}

func (circle Circle) area() float64 {
	return 3.14 * circle.radius * circle.radius
}

// 為了提高效率，通常我們方法和結構體指針類型綁定
func (circle *Circle) area2() float64 {
	//因為c是指針，因此我們標準的訪問其字段的方式是(*circle).radius
	//(*circle).radius 等價 circle.radius
	fmt.Printf("circle指向的地址：%p \n", circle)
	circle.radius = 10
	return 3.14 * (*circle).radius * (*circle).radius
}

func main() {
	//1. 聲明一個結構體`Circle`，字段為`radius`。
	//2. 聲明一個方法`area`和`Circle`綁定，可以返回面積。

	circle := Circle{4.0}
	res := circle.area()
	fmt.Println("面積為：", res)

	fmt.Println()

	circle2 := Circle{7.0}
	fmt.Printf("main circle2結構體變數地址=%p \n", &circle2)

	//res2 := (&circle2).area2()
	//編譯器底層做了優化， (&circle).area2() 等價 circle.area2()
	//編譯器會自動幫我們加上&
	res2 := circle2.area2()

	fmt.Println("circle2 radius=", circle2.radius)
	fmt.Println("面積2為：", res2)
}
