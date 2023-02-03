package main

import "fmt"

func printNineNineMultiplicationTable(num int) {
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d \t", j, i, j*i)
		}
		fmt.Println()
	}
}

func main() {
	printNineNineMultiplicationTable(9)
}
