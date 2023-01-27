package main

import "fmt"

func main() {
	//編寫程序，聲明2個`int32`變數並賦值，判斷兩數之和，如果大於等於50，打印`hello world!`。
	var n1 int32 = 25
	var n2 int32 = 26

	if n1+n2 >= 50 {
		fmt.Println("hello world!")
	}

	//編寫程序，聲明2個`float64`變數並賦值，判斷第一個數大於`10.0`且第二個數小於`20.0`，打印兩數之和。
	var f1 float64 = 11.0
	var f2 float64 = 15.0
	if f1 > 10.0 && f2 < 20.0 {
		fmt.Println(f1 + f2)
	}

	//定義兩個`int32`變數，判斷兩者的和，是否能被3和5整除，打印提示信息。
	var n3 int32 = 30
	var n4 int32 = 60
	if (n3+n4)%15 == 0 {
		fmt.Println("可以被3和5整除")
	} else {
		fmt.Println("不能被3和5整除")
	}

	//判斷一個年份是否為閏年，閏年的條件是符合下面兩者之一：(1)年份能被4整除但不能被100整除 (2)能被400整除。
	var year = 2020
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Println(year, "是閏年")
	} else {
		fmt.Println(year, "不是閏年")
	}

}
