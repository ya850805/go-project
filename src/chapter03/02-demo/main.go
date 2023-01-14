package main

import "fmt"

func main() {
	//1. 第一種：指定變數類型，聲明後若不賦值，使用默認值。
	//int的默認值是0
	var i int
	fmt.Println("i=", i)

	//第二種：根據值自行判定變數類型(**類型推導**)。
	var num = 10.11
	fmt.Println("num=", num)

	//第三種：省略`var`，:warning:`:=`左側的變數不應該是已經聲明過的，否則會導致編譯錯誤。
	//下面的方式等價於 var name = "jason"
	//:=的:不能省略，否則編譯錯誤
	name := "jason"
	fmt.Println("name=", name)
}
