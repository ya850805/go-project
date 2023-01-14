package main

import "fmt"

// 定義全局變數
var globalN1 = 100
var globalN2 = 200
var globalName1 = "juan1"

// 上面的聲明方式，也可以改成一次性聲明
var (
	globalN3    = 300
	globalN4    = 400
	globalName2 = "juan2"
)

func main() {
	//golang如何一次性聲明多個變數
	var n1, n2, n3 int
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)

	//一次性聲明多個不同類型的變數
	var num1, name1, num2 = 100, "jason", 200
	fmt.Println("num1=", num1, "name1=", name1, "num2=", num2)

	//一次性聲明多個不同類型的變數(省略var)
	num3, name2, num4 := 300, "jason2", 400
	fmt.Println("num3=", num3, "name2=", name2, "num4=", num4)

	//輸出全局變數
	fmt.Println("globalN1=", globalN1, "globalN2=", globalN2, "globalName1=", globalName1)
	fmt.Println("globalN3=", globalN3, "globalN4=", globalN4, "globalName2=", globalName2)
}
