package main

import (
	"fmt"
)

func putNum(intChan chan int) {
	for i := 1; i <= 80000; i++ {
		intChan <- i
	}

	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	for {
		num, ok := <-intChan
		if !ok {
			break
		}

		isPrime := true
		//查看是不是質數
		for i := 2; i <= num-1; i++ {
			if num%i == 0 { //說明num不是質數
				isPrime = false
				break
			}
		}

		if isPrime {
			primeChan <- num
		}
	}

	fmt.Println("有一個primeNum協程因為取不到數據而退出")
	//這邊我們還不能關閉primeChan(因為可能還有其他primeNum協程在工作)
	exitChan <- true
}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000) //放入的結果(質數)
	exitChan := make(chan bool, 4)    //標示退出的管道

	//開啟一個協程，向intChan放入1-8000的數
	go putNum(intChan)

	//開啟4個協程，從intChan取數據，並判斷是不是質數，如果是就放入primeChan
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	//從exit取出4個true，當可以取完4個時，代表全部完成
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}

		close(primeChan)
	}()

	//遍歷primeChan取出結果
	for {
		prime, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println("取出質數：", prime)
	}
}
