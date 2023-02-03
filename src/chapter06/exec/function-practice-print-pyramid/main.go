package main

import "fmt"

func printPyramid(layers int) {
	for i := 1; i <= layers; i++ {
		//印空格
		for j := 1; j <= layers-i; j++ {
			fmt.Print(" ")
		}

		//印星號
		for j := 1; j <= (2*i)-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	var layers int
	fmt.Print("請輸入層數：")
	fmt.Scanln(&layers)

	printPyramid(layers)
}
