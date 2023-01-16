package main

import "fmt"

func main() {
	//在Golang中，++和--只能獨立使用
	var i int = 8
	var a int
	//報錯，因為i++, i--只能獨立使用
	//a = i++
	//a = i--

	//if i++ > 0 {
	//
	//}
	fmt.Println("i=", i)
	fmt.Println("a=", a)

	var num int = i
	num++
	//報錯，Golang中沒有前++，也沒有前--
	//++num
}
