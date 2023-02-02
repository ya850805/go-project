package main

import "fmt"

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	//在Go中，**函數也是一種數據類型**，可以賦值給一個變數，則該變數就是一個函數類型的變數了。通過該變數可以對函數調用
	a := getSum
	fmt.Printf("a的類型是%T, getSum的類型是%T \n", a, getSum)

	res := a(10, 40) //等價getSum(10, 40)
	fmt.Println("res=", res)
}
