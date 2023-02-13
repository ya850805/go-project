package main

import "fmt"

type A struct {
	Num int
}

type B struct {
	Num int
}

func main() {
	var a A
	var b B

	//結構體是用戶單獨定義的類型，和其他類型進行轉換時需要有完全相同的字段名(名稱、個數和類型)
	//可以轉換，可是結構體的字段要完全一樣
	a = A(b)
	fmt.Println("a=", a)
	fmt.Println("b=", b)
}
