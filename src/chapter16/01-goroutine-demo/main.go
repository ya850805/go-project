package main

import (
	"fmt"
	"strconv"
	"time"
)

//1. 在主線程(可以理解為進程)中，開啟一個goroutine，該協程每隔一秒輸出"hello,world"。
//2. 在主線程中也每隔一秒輸出"hello,golang"，輸出10次後，退出程序。
//3. 要求主線程和goroutine同時執行。

// 編寫一個函數，每隔一秒輸出"hello,world"
func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("test() hello,world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	go test() //開啟一個協程

	for i := 1; i <= 10; i++ {
		fmt.Println("main() hello,golang" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
