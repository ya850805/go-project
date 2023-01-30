package main

import "fmt"

func main() {
	var sum = 6
	for i := 0; i <= sum; i++ {
		fmt.Printf("%v + %v = %v \n", i, sum-i, sum)
	}
}
