package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) test() {
	p.Name = "Juan"
	fmt.Println("test() name=", p.Name)
}

// 給`Person`結構體添加一個`speak`方法，輸出`xxx是一個好人`。
func (p Person) speak() {
	fmt.Println(p.Name, "is a good man")
}

// 給`Person`結構體添加計算方法，可以計算從1+...+1000的結果。
func (p Person) cal() {
	sum := 0
	for i := 1; i <= 1000; i++ {
		sum += i
	}
	fmt.Println(p.Name, "計算的結果是：", sum)
}

// 給`Person`結構體添加一個計算方法，該方法可以接收一個數`n`，計算1+...+n的結果。
func (p Person) cal2(n int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Println(p.Name, "計算的結果是：", sum)
}

// 給`Person`結構體添加`getSum`方法，可以計算兩個數的和，並返回結果。
func (p Person) getSum(n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	var p Person
	p.Name = "Jason"
	p.test()
	fmt.Println("p.Name=", p.Name)

	p.speak()
	p.cal()
	p.cal2(10)
	n1 := 10
	n2 := 20
	sum := p.getSum(n1, n2)
	fmt.Println("sum=", sum)
}
