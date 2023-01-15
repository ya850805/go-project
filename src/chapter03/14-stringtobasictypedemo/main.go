package main

import (
	"fmt"
	"strconv"
)

// golang中string轉基本數據類型
func main() {
	var str string = "true"
	//`strconv.ParseBool`函數會返回兩個值(value bool, err error)
	//因為我只想獲取到value bool，不想獲取err，所以使用_忽略
	b, _ := strconv.ParseBool(str)
	fmt.Printf("b type is %T, b=%v \n", b, b)

	var str2 string = "123456789"
	n1, _ := strconv.ParseInt(str2, 10, 64)
	fmt.Printf("n1 type is %T, n1=%v \n", n1, n1)
	n2 := int(n1)
	fmt.Printf("n2 type is %T, n1=%v \n", n2, n2)

	var str3 string = "123.456"
	f1, _ := strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type is %T, n1=%v \n", f1, f1)

	//注意事項，轉成int失敗時默認會轉成0，轉成bool失敗時默認會轉成false
	var str4 string = "hello"
	n3, _ := strconv.ParseInt(str4, 10, 64)
	fmt.Printf("n3 type is %T, n3=%v \n", n3, n3)

	var str5 string = "hello"
	b1, _ := strconv.ParseBool(str5)
	fmt.Printf("b1 type is %T, b1=%v \n", b1, b1)
}
