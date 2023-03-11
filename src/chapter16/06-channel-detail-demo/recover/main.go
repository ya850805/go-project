package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello world")
	}
}

func test() {
	defer func() {
		//捕獲test拋出的panic
		if err := recover(); err != nil {
			fmt.Println("test() 發生錯誤：", err)
		}
	}()

	var myMap map[int]string
	//使用map沒有make，會有panic
	myMap[0] = "golang"
}

func main() {
	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("main() ===", i)
	}
}
