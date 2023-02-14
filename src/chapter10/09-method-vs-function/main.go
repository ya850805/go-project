package main

import "fmt"

type Person struct {
	Name string
}

// 函數
// 對於普通**函數**，**接收者為值類型時，不能將指針類型的數據直接傳遞**，反之亦然
func test01(p Person) {
	fmt.Println("test01()=", p.Name)
}

// 方法
// 對於**方法**(如`struct`的方法)，**接收者為值類型時，可以值接用指針類型的變數調用方法**，反之亦然
func (p Person) test02() {
	fmt.Println("test02()=", p.Name)
}

func (p *Person) test03() {
	p.Name = "marry"
	fmt.Println("test03()=", p.Name)
}

// 1. **不管調用形式如何，真正決定是值拷貝還是地址拷貝，是看這個方法和哪個類型綁定**。
// 2. 如果方法是和值類型綁定(`(p Person)`)，則是值拷貝，如果是和指針類型綁定(`(p *Person)`)，則是地址拷貝。
func main() {
	p := Person{"tom"}
	test01(p)
	p.test02()
	(&p).test02() // 從形式上是傳入地址，但本質仍然是值拷貝(主要看該方法是跟什麼綁定的)

	//是傳地址
	p.test03()
	(&p).test03()
}
