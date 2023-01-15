package main

import (
	"fmt"
	"strconv"
)

// golang基本數據類型轉string
func main() {
	//第一種方式：`fmt.Sprintf()`
	var num1 int = 10
	var num2 float64 = 23.456
	var b bool = true
	var myChar byte = 'h'

	var str string //空的字符串

	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type is %T, str=%q \n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type is %T, str=%q \n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type is %T, str=%q \n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str type is %T, str=%q \n", str, str)

	/////////////////////////////////////////////////////////////////////////////////////

	//第二種方式：使用`strconv`包的函數
	var num3 int = 99
	var num4 float64 = 66.666
	var b2 bool = true

	var str2 string //空的字符串

	str2 = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str2 type is %T, str2=%q \n", str2, str2)

	//說明：'f'：格式, 10：表示小數位保留10位, 64：表示這個小數是float64
	str2 = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str2 type is %T, str2=%q \n", str2, str2)

	str2 = strconv.FormatBool(b2)
	fmt.Printf("str2 type is %T, str2=%q \n", str2, str2)

	//strconv包中有一個函數Itoa，將一個int轉換成string
	var num5 int = 999
	str2 = strconv.Itoa(num5)
	fmt.Printf("str2 type is %T, str2=%q \n", str2, str2)
}
