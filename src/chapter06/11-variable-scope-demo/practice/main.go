package main

import "fmt"

var name = "tom"

//這裡會報錯，此賦值語句不能在外邊執行
//Name := "mary"

func test1() {
	fmt.Println(name)
}

func test2() {
	name := "jack"
	fmt.Println(name)
}

func main() {
	fmt.Println(name)
	test1()
	test2()
	test1()
}
