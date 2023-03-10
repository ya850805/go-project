package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {

	//定義一個可以存放任意數據的管道，長度為3
	var allChan chan interface{}
	allChan = make(chan interface{}, 3)

	allChan <- 10
	allChan <- "tom"
	cat := Cat{Name: "cat1", Age: 10}
	allChan <- cat

	//為了取到Cat物件，先推出前2個
	<-allChan
	<-allChan

	newCat := <-allChan //從管道取出的Cat

	//編譯時，newCat是一個interface{}類型，如果要取他的屬性需要類型斷言
	name := newCat.(Cat).Name
	fmt.Println("name=", name)

	fmt.Printf("newCat類型是%T, newCat=%v", newCat, newCat)
}
