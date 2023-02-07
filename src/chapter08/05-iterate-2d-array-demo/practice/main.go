package main

import "fmt"

func main() {
	//定義二維陣列，用於保存三個班，每班五名同學成績，
	//並求出每個班級平均分以及所有班級平均分。

	//1. 定義一個二維陣列
	var scores [3][5]float64

	//2. 循環輸入成績
	for i := 0; i < len(scores); i++ {
		for j := 0; j < len(scores[i]); j++ {
			fmt.Printf("請輸入第%d個班 第%d名學生的成績：", i+1, j+1)
			fmt.Scanln(&scores[i][j])
		}
	}

	//3. 遍歷輸出成績後的二維陣列，統計平均分
	totalSum := 0.0
	for i := 0; i < len(scores); i++ {
		sum := 0.0
		for j := 0; j < len(scores[i]); j++ {
			sum += scores[i][j]
		}
		totalSum += sum
		fmt.Printf("第%d個班的平均分為：%v \n", i+1, sum/float64(len(scores[i])))
	}

	fmt.Println("所有班級的平均分為：", totalSum/float64(len(scores)*len(scores[0])))
}
