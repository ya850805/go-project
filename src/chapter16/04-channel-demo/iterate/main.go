package main

import "fmt"

func main() {
	intChan := make(chan int, 100)

	//放入100個數到管道
	for i := 0; i < 100; i++ {
		intChan <- i * 2
	}

	//在遍歷時，如果channel沒有關閉，則會出現deadlock的錯誤
	//因此我們先關閉管道
	close(intChan)

	//使用for-range遍歷
	//在遍歷時，如果channel已經關閉，則會正常遍歷數據，遍歷完後，就會退出遍歷
	for v := range intChan {
		fmt.Println("v=", v)
	}
}
