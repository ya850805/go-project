package main

import "fmt"

func sum(n1 int, n2 int) int {
	//當執行到defer時，暫時不執行，會將defer後面的語句壓入到獨立的棧中(defer棧)
	//當函數執行完畢後，再從defer棧，按照"先入後出"的方式出棧，並執行
	defer fmt.Println("ok1 n1=", n1)
	defer fmt.Println("ok2 n2=", n2)

	res := n1 + n2
	fmt.Println("ok3 res=", res)
	return res
}

func main() {
	res := sum(10, 20)
	fmt.Println("res=", res)
}
