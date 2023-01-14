package main

import "fmt"

// golang小數類型使用
func main() {
	var price float32 = 89.12
	fmt.Println("price=", price)

	var num1 float32 = -0.00089
	var num2 float64 = -7809656.09
	fmt.Println("num1=", num1, "num2=", num2)

	//尾數部分可能會丟失，造成精度損失。
	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901
	fmt.Println("num3=", num3, "num4=", num4)

	//Golang浮點型默認聲明為float64類型
	var num5 = 1.1
	fmt.Printf("num5的類型是%T \n", num5)

	//十進制數形式：如5.12, .512(必須有小數點)
	num6 := 5.12
	num7 := .512 //0.512
	fmt.Println("num6=", num6, "num7=", num7)

	//科學計數形式
	num8 := 5.1234e2   //5.1234 * 10的2次方
	num9 := 5.1234e2   //5.1234 * 10的2次方
	num10 := 5.1234e-2 //5.1234 / 10的2次方
	fmt.Println("num8=", num8, "num9=", num9, "num10=", num10)
}
