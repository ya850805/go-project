package main

import "fmt"

func main() {
	//演示/ 、 %

	//如果參與運算的都是整數，那麼除後會去掉小數部分，保留整數部分
	fmt.Println("10 / 4 =", 10/4)
	var n1 float32 = 10 / 4
	fmt.Println("n1=", n1)

	//如果希望保留小數部分，就需要有浮點數參與運算
	var n2 float32 = 10.0 / 4
	fmt.Println("n2=", n2)

	//演示%
	//看一個公式 a % b = a - ((a / b) * b)
	fmt.Println("10 % 3 =", 10%3)
	fmt.Println("-10 % 3 =", -10%3)
	fmt.Println("10 % -3 =", 10%-3)
	fmt.Println("-10 % -3 =", -10%-3)

	//++和--的使用
	var i int = 10
	i++ //等價 i = i + 1
	fmt.Println("i=", i)
	i-- //等價 i = i - 1
	fmt.Println("i=", i)
}
