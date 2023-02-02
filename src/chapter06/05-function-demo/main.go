package main

import "fmt"

// another example, rename function type
type myFunType func(int, int) int

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

// 函數既然是一種數據類型，因此在Go中，函數可以做為參數，並且調用。
func myFunc(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}

func myFunc2(funvar myFunType, num1 int, num2 int) int {
	return funvar(num1, num2)
}

func main() {
	//在Go中，**函數也是一種數據類型**，可以賦值給一個變數，則該變數就是一個函數類型的變數了。通過該變數可以對函數調用
	a := getSum
	fmt.Printf("a的類型是%T, getSum的類型是%T \n", a, getSum)

	res := a(10, 40) //等價getSum(10, 40)
	fmt.Println("res=", res)

	res2 := myFunc(getSum, 50, 60)
	fmt.Println("res2=", res2)

	//給int取別名，在go中myInt和int雖然都是int類型，但是go認為myInt和int兩個是不同類型
	type myInt int
	var num1 myInt

	var num2 int
	num2 = int(num1)
	fmt.Printf("num1=%v, num1類型是: %T \n", num1, num1)
	fmt.Printf("num2=%v, num2類型是: %T \n", num2, num2)

}
