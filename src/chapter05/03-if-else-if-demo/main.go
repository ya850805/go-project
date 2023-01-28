package main

import "fmt"

func main() {
	//小明參加考試，他與父親達成承諾：
	//如果
	//* 成績為100分，獎勵一輛BMW。
	//* 成績為(80,99]時，獎勵一台iphone。
	//* 當成績為[60,80]時，獎勵一台ipad。
	//* 其他成績什麼獎勵都沒有。
	//
	//請從鍵盤輸入小明的成績，並加以判斷

	var score int
	fmt.Print("請輸入小明的成績：")
	fmt.Scanln(&score)

	if score == 100 {
		fmt.Println("獎勵一輛BMW")
	} else if score > 80 && score <= 99 {
		fmt.Println("獎勵一台iphone")
	} else if score >= 60 && score <= 80 {
		fmt.Println("獎勵一台ipad")
	} else {
		fmt.Println("什麼都不獎勵")
	}

	//使用陷阱
	//下面代碼只會輸出ok1
	var n int = 10
	if n > 9 {
		fmt.Println("ok1")
	} else if n > 6 {
		fmt.Println("ok2")
	} else if n > 3 {
		fmt.Println("ok3")
	} else {
		fmt.Println("ok4")
	}
}
