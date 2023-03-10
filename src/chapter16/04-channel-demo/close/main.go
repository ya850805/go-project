package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200

	//關閉管道
	close(intChan)

	n1 := <-intChan
	fmt.Println("n1=", n1)

	//報錯，panic: send on closed channel
	intChan <- 300
}
