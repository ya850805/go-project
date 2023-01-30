package main

import "fmt"

func main() {
	//打印1~100之間所有是9的倍數的整數的個數及總和
	var count, sum = 0, 0
	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			count++
			sum += i
		}
	}
	fmt.Printf("1~100中9的倍數總共有%d個，其總和是%d", count, sum)
}
