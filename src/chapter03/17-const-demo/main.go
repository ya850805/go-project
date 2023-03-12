package main

import "fmt"

func main() {
	var num int
	num = 90 //ok

	//常數聲明時，必須賦值
	const tax int = 0
	//常數不能修改
	//tax = 10
	fmt.Println("num=", num, "tax=", tax)

	//一行加一次1
	const (
		a = iota
		b
		c
		d
	)

	fmt.Println(a, b, c, d)
}
