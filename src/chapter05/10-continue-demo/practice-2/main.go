package main

import "fmt"

func main() {
	//某人有100000元，每經過一次路口需要繳費，規則如下
	//當現金>50000時，每次繳5%
	//當現金<=50000時，每次繳1000
	//計算該人可以經過多少次路口，使用`for break`完成

	time := 0
	money := 100000.0
	for {
		if money < 1000 {
			break
		}

		if money > 50000 {
			money = money * 0.95
		} else {
			money -= 1000
		}
		time++
	}

	fmt.Printf("經過%d次路口", time)
}
