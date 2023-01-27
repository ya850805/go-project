package main

import "fmt"

func main() {
	//對下方的代碼，如果有輸出，指出輸出的結果
	var i int = 4
	var j int = 1
	if i > 2 {
		if j > 2 {
			fmt.Println(i + j)
		}
		fmt.Println("hello!")
	} else {
		fmt.Println("i=", i)
	}

	//輸出結果是hello
}
