package main

import "fmt"

func main() {
	//`continue`實現打印1~100之內的奇數(要求使用`for`循環+`continue`)
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	//從鍵盤輸入個數不確定的**整數**，並判斷讀入的正數和負數的個數，輸入為0時結束程序(使用`for`循環、`break`、`continue`完成)
	var num int
	var positiveCount = 0
	var negativeCount = 0
	for {
		fmt.Print("請輸入整數：")
		fmt.Scanln(&num)

		if num == 0 {
			break
		}

		if num > 0 {
			positiveCount++
			continue
		}
		negativeCount++
	}
	fmt.Printf("正數有%d個，負數有%d個", positiveCount, negativeCount)
}
