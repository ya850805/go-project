package main

import "fmt"

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
}
