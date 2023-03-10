package main

import "fmt"

func main() {
	//管道的使用

	//1. 創建一個可以存放3個int類型的管道
	var intChan chan int
	intChan = make(chan int, 3)

	//2. 看看intChan是什麼
	fmt.Printf("intChan的值=%v intChan的地址=%p \n", intChan, &intChan)

	//3. 向管道寫入數據
	intChan <- 10
	num := 211
	intChan <- num

	//4. 看看管道的長度和cap(容量)
	fmt.Printf("intChan len=%v cap=%v \n", len(intChan), cap(intChan))

	//注意點：當我們給管道寫入數據時，不能超過其容量
	intChan <- 3
	//intChan <- 4
	fmt.Printf("intChan len=%v cap=%v \n", len(intChan), cap(intChan))

	//5. 從管道中讀取數據
	num2 := <-intChan
	fmt.Println("num2=", num2)
	fmt.Printf("intChan len=%v cap=%v \n", len(intChan), cap(intChan))

	num3 := <-intChan
	num4 := <-intChan
	fmt.Println("num3=", num3)
	fmt.Println("num4=", num4)

	//6. 在沒有使用協程的情況下，如果我們的管道數據已經全部取出，再去取就會報告deadlock
	num5 := <-intChan
	fmt.Println("num5=", num5)

}
