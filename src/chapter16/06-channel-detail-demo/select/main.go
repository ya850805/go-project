package main

import "fmt"

func main() {
	//使用select可以解決從管道取數據的阻塞問題

	//1. 定義一個管道10個int數據
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	//2. 定義一個管道5個string數據
	strChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		strChan <- "hello" + fmt.Sprintf("%d", i)
	}

	//傳統的方法在遍歷管道時，如果不關閉管道會阻塞而導致deadlock

	//問題：在實際開發中，可能我們不好確定什麼時候關閉該管道
label:
	for {
		select {
		//注意：如果這邊intChan管道一直沒有關閉，不會一直阻塞而deadlock
		//會自動到下一個case匹配
		case v := <-intChan:
			fmt.Printf("從intChan讀取了數據%d \n", v)
		case v := <-strChan:
			fmt.Printf("從strChan讀取了數據%s \n", v)
		default:
			fmt.Println("intChan和strChan都取不到了")
			break label
		}
	}
}
