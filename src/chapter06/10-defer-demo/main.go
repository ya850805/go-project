package main

import "fmt"

func sum(n1 int, n2 int) int {
	//當執行到defer時，暫時不執行，會將defer後面的語句壓入到獨立的棧中(defer棧)
	//當函數執行完畢後，再從defer棧，按照"先入後出"的方式出棧，並執行
	//在defer將語句放入到棧時，也會將相關的值拷貝同時入棧
	defer fmt.Println("ok1 n1=", n1) //也將n1做值拷貝進defer棧
	defer fmt.Println("ok2 n2=", n2) //也將n2做值拷貝進defer棧

	//增加一句話
	n1++ //n1=11
	n2++ //n2=21

	res := n1 + n2
	fmt.Println("ok3 res=", res)
	return res
}

func main() {
	res := sum(10, 20)
	fmt.Println("res=", res)
}
