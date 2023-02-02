package main

import "fmt"

func main() {
	n1 := 10
	//調用test()
	test(n1)
	fmt.Println("main函數n1=", n1)

	//調用getSum()
	sum := getSum(10, 20)
	fmt.Println("main sum=", sum)

	//調用getSumAndSub()
	res1, res2 := getSumAndSub(1, 2)
	fmt.Printf("res1=%d, res2=%d \n", res1, res2)

	//希望忽略某個返回值，則使用_符號表示佔位忽略
	_, res3 := getSumAndSub(31, 9)
	fmt.Printf("res3=%d", res3)
}

func test(n1 int) {
	n1 = n1 + 1
	fmt.Println("test函數n1=", n1)
}

func getSum(n1 int, n2 int) int {
	sum := n1 + n2
	fmt.Println("getSum sum=", sum)
	//當函數有return語句時，就是將結果返回給調用者
	//即誰調用我，就返回給誰
	return sum
}

func getSumAndSub(n1 int, n2 int) (int, int) {
	return n1 + n2, n1 - n2
}
