package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}

	a = point //OK

	//如何將a賦給一個Point變數？使用類型斷言
	var point2 Point
	point2 = a.(Point)
	fmt.Println(point2)

	//類型斷言的其他案例
	//var x interface{}
	//var b float32 = 1.1
	//x = b //空介面可以接收任何類型
	//y := x.(float32)
	//fmt.Printf("y的類型是%T 值是%v", y, y)

	//帶檢測的類型斷言
	var x interface{}
	var b float32 = 1.1
	x = b //空介面可以接收任何類型
	y, ok := x.(float64)
	if ok {
		fmt.Println("轉換成功")
		fmt.Printf("y的類型是%T 值是%v", y, y)
	} else {
		fmt.Println("轉換失敗")
	}

}
