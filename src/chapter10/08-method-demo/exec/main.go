package main

import "fmt"

type MethodUtils struct {
}

type Calculator struct {
	Num1 float64
	Num2 float64
}

func (mu *MethodUtils) Print() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (mu *MethodUtils) Print2(m int, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (mu *MethodUtils) Area(len float64, width float64) float64 {
	return len * width
}

func (mu *MethodUtils) JudgeNum(num int) {
	if num%2 == 0 {
		fmt.Println(num, "是偶數")
	} else {
		fmt.Println(num, "是奇數")
	}
}

func (mu *MethodUtils) Print3(n int, m int, key string) {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(key)
		}
		fmt.Println()
	}
}

func (c *Calculator) cal(operator byte) float64 {
	var res float64
	n1 := c.Num1
	n2 := c.Num2
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("運算符輸入有誤...")
	}
	return res
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
	area := methodUtils.Area(10, 20)
	fmt.Println("面積為：", area)

	fmt.Println()

	//4. 編寫方法：判斷一個數是奇數還是偶數
	methodUtils.JudgeNum(14)
	methodUtils.JudgeNum(11)

	fmt.Println()

	//5. 根據行、列、字符打印對應行數和列數的字符，例如：行:3、列:2、字符:`*`，則打印相應的效果
	methodUtils.Print3(3, 2, "#")

	fmt.Println()

	//定義一個計算器結構體，實現加減乘除四個功能，用一個方法完成，傳入一個運算符，結構體有2個數的屬性。
	cal := Calculator{10, 5}
	res := cal.cal('*')
	fmt.Println("計算的結果：", res)
}
