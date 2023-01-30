package main

import "fmt"

func main() {
	//統計3個班成績情況，每個班有5名同學，求出各個班的平均分數和所有班級的平均分(學生成績從鍵盤輸入)。
	//統計3個班的及格人數，每個班有5名同學。

	var totalSum = 0.0
	var passCount = 0
	for i := 1; i <= 3; i++ {
		var classSum = 0.0
		for j := 1; j <= 5; j++ {
			var score float64
			fmt.Printf("請輸入%d個班，第%d個同學的分數：", i, j)
			fmt.Scanln(&score)
			if score >= 60 {
				passCount++
			}
			classSum += score
		}
		fmt.Printf("第%d個班的平均分數是：%v \n", i, classSum/5)
		totalSum += classSum
	}
	fmt.Println("所有同學的平均分數是：", totalSum/15)
	fmt.Println("及格人數是：", passCount)
}
