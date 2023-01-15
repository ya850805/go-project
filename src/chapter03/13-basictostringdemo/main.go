package main

import "fmt"

// golang基本數據類型轉string
func main() {
	var num1 int = 10
	var num2 float64 = 23.456
	var b bool = true
	var myChar byte = 'h'

	var str string //空的字符串

	//第一種方式：`fmt.Sprintf()`
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type is %T, str=%q \n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type is %T, str=%q \n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type is %T, str=%q \n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str type is %T, str=%q \n", str, str)
}
