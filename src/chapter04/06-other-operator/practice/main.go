package main

import "fmt"

func main() {
	//求兩個數的最大值
	var i = 10
	var j = 20
	if i > j {
		fmt.Println("i,j的最大值為=", i)
	} else {
		fmt.Println("i,j的最大值為=", j)
	}
	//求三個數的最大值
	//思路：先求出兩個數的最大值，然後讓這個最大值和第三數比較，再取出最大值
	var n1 = 10
	var n2 = 20
	var n3 = 30

	var max int
	if n1 > n2 {
		max = n1
	} else {
		max = n2
	}

	if n3 > max {
		max = n3
	}
	fmt.Println("三個數的最大值=", max)

}
