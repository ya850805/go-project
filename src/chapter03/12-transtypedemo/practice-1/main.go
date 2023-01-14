package main

import "fmt"

func main() {
	//練習題1
	var i1 int32 = 12
	var i2 int64
	var i3 int8

	//Cannot use 'i1 + 20' (type int32) as the type int64
	//i2 = i1 + 20
	//Cannot use 'i1 + 20' (type int32) as the type int
	//i3 = i1 + 20

	//修正
	i2 = int64(i1 + 20)
	i3 = int8(i1 + 20)
	fmt.Println("i2=", i2, "i3=", i3)
}
