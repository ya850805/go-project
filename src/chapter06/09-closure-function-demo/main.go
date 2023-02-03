package main

import "fmt"

//golang中的閉包

// AddUpper 累加器
func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n += x
		return n
	}
}

func main() {
	f := AddUpper()
	fmt.Println(f(1)) //11
	fmt.Println(f(2)) //13
	fmt.Println(f(3)) //16
}
