package main

import "fmt"

func main() {
	//使用遞歸的方式，求出費波那契數1,1,2,3,5,8,13...。給一個整數n，求出他的費波那契數是多少？
	n := 6
	res := getFibonacci(n)

	fmt.Println("res=", res)
}

func getFibonacci(n int) int {
	if n <= 2 {
		return 1
	} else {
		return getFibonacci(n-1) + getFibonacci(n-2)
	}
}
