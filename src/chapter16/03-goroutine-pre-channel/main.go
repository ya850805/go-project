package main

import (
	"fmt"
	"sync"
	"time"
)

//現在要計算1-200的各個數的階乘，並且把各個數的階乘放到map中。最後顯示出來。要求使用goroutine完成

//思路
//1. 編寫一個函數，來計算各個數的階乘，並放到map中
//2. 我們啟動多個協程，統計的結果放到map中
//3. map應該做成全局的

var (
	myMap = make(map[int]int)

	//聲明一個全局的互斥鎖
	//lock是一個全局的互斥鎖
	//sync是包，synchronized同步
	//Mutex：互斥
	lock sync.Mutex
)

// cal函數就是計算n!，然後將結果放到myMap
func cal(n int) {
	res := 1

	for i := 1; i <= n; i++ {
		res *= i
	}

	//加鎖
	lock.Lock()

	myMap[n] = res

	//解鎖
	lock.Unlock()
}

func main() {
	for i := 1; i <= 20; i++ {
		go cal(i)
	}

	time.Sleep(time.Second * 10)

	lock.Lock()
	for i1, i2 := range myMap {
		fmt.Printf("%v的階乘是%v \n", i1, i2)
	}
	lock.Unlock()
}
