package main

import "fmt"

func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		//放入數據
		intChan <- i
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan

		if !ok {
			break
		}

		fmt.Printf("readData 讀到數據%v \n", v)
	}

	//讀完所有數據後，即任務完成
	exitChan <- true
	close(exitChan)
}

func main() {
	//創建兩個管道
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	//讀退出的管道，當退出的管道關閉了，代表readData讀完全部數據了
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
