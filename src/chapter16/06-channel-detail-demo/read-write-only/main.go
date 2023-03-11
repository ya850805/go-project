package main

import "fmt"

// 在方法中聲明為只寫，避免誤用
func send(intChan chan<- int) {

}

// 在方法中聲明為只讀，避免誤用
func receive(intChan <-chan int) {

}

func main() {
	//管道可以聲明為只讀或只寫

	//1. 默認情況下，管道是雙向(可寫可讀)
	var chan1 chan int //可讀可寫
	fmt.Println("chan1=", chan1)

	//2. 聲明為只寫
	var chan2 chan<- int
	chan2 = make(chan int, 3)
	chan2 <- 20

	//報錯，chan2只能寫
	//num := <- chan2
	fmt.Println("chan2=", chan2)

	//3. 聲明為只讀
	var chan3 <-chan int
	//報錯，chan3只能讀
	//chan3 <- 10
	fmt.Println("chan3=", chan3)

	//使用案例
	var intChan chan int //雙向
	intChan = make(chan int, 3)

	send(intChan)
	receive(intChan)
}
