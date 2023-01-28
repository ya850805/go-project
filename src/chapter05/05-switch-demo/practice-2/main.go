package main

import "fmt"

func main() {
	//對學生成績大於60分的，輸出"合格"，低於60分的，輸出"不合格"。(輸入的成績不能大於100)
	var score float64
	fmt.Print("請輸入成績：")
	fmt.Scanln(&score)

	switch {
	case score >= 60 && score <= 100:
		fmt.Println("合格")
	case score < 60 && score >= 0:
		fmt.Println("不合格")
	default:
		fmt.Println("輸入成績有誤")
	}
}
