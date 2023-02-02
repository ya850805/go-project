package main

import "fmt"

func main() {
	//已知`f(1)=3` `f(n) = 2 * f(n - 1) + 1`
	//請用遞歸求出`f(n)`的值。

	n := 5
	res := getNum(n)

	fmt.Println("res=", res)
}

func getNum(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*getNum(n-1) + 1
	}
}
