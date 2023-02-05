package main

import "fmt"

func main() {
	//從終端循環輸入5個成績，保存到`float64`陣列，並輸出
	var scoreArray [5]float64
	for i := 0; i < len(scoreArray); i++ {
		fmt.Printf("請輸入第%d個成績：", i+1)
		fmt.Scanln(&scoreArray[i])
	}

	fmt.Println("成績陣列為=", scoreArray)
}
