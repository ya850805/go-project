package main

import "fmt"

type MethodUtils struct {
}

func (mu MethodUtils) Print() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (mu MethodUtils) Print2(m int, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (mu MethodUtils) area(len float64, width float64) float64 {
	return len * width
}

func main() {
	//1. 編寫結構體(`MethodUtils`)，編寫一個方法，方法不需要參數，在方法中打印一個10\*8的矩形，在`main()`方法中調用該方法
	var methodUtils MethodUtils
	methodUtils.Print()

	fmt.Println()

	//2. 編寫一個方法，提供`m`和`n`兩個參數，方法中打印一個`m`\*`n`的矩形。
	methodUtils.Print2(2, 6)

	fmt.Println()

	//3.編寫一個方法算矩形的面積(可以接收長`len`和寬`width`)，將其作為方法返回值。在`main()`方法中調用該方法，接收返回的面積值並打印。
	area := methodUtils.area(10, 20)
	fmt.Println("面積為：", area)
}
