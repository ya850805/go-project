package main

import "fmt"

func main() {
	var layers int
	fmt.Print("請輸入層數：")
	fmt.Scanln(&layers)

	//第一種
	for i := 1; i <= layers; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println("======================================================")

	//第二種
	for i := 1; i <= layers; i++ {
		//打印空格，規律：總層數-當前層數
		for j := 1; j <= layers-i; j++ {
			fmt.Print(" ")
		}

		//打印*，規律：2 * 層數 - 1
		for k := 1; k <= i*2-1; k++ {
			fmt.Print("*")
		}

		fmt.Println()
	}

	fmt.Println("======================================================")

	//第三種：空心金字塔，在我們給每行打印*時，需要考慮是打印*還是空格
	for i := 1; i <= layers; i++ {
		//打印空格，規律：總層數-當前層數
		for j := 1; j <= layers-i; j++ {
			fmt.Print(" ")
		}

		//打印*，規律：2 * 層數 - 1
		for k := 1; k <= i*2-1; k++ {
			//每行的第一個和最後一個 或者 底層全部打星
			if k == 1 || k == i*2-1 || i == layers {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}
}
