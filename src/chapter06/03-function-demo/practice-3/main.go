package main

import "fmt"

func main() {
	//有一堆桃子，猴子第一天吃了其中的一半，並再多吃了一個。
	//之後每天猴子都吃了其中的一半且多一個。
	//到第10天時，想再吃時發現只剩一個桃子。請問最初多少個桃子？

	res := getPeach(1)
	fmt.Printf("一開始有%d個桃子", res)
}

func getPeach(day int) int {
	if day == 10 {
		return 1
	} else {
		return (getPeach(day+1) + 1) * 2
	}
}
