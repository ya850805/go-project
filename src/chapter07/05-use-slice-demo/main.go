package main

import "fmt"

func main() {
	//演示切片的使用make
	var slice []float64 = make([]float64, 5, 10)
	slice[1] = 10
	slice[2] = 20
	fmt.Println("slice=", slice)
	fmt.Println("slice的size=", len(slice))
	fmt.Println("slice的cap=", cap(slice))

	//定義一個切片，直接就指定具體陣列，使用原理類似`make`的方式
	var slice2 = []int{1, 3, 5}
	fmt.Println("slice2=", slice2)
}
