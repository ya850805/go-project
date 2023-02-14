package main

import "fmt"

type integer int

func (i *integer) print() {
	fmt.Println("i=", i)
}

type Student struct {
	Name string
	Age  int
}

// 如果一個類型實現了`String()`這個方法，那麼`fmt.Println`默認會默認會調用這個變數的`String()`進行輸出。
func (s *Student) String() string {
	str := fmt.Sprintf("~~Name=[%v], Age=[%v]~~", s.Name, s.Age)
	return str
}

// 編寫一個方法，可以改變i的值
func (i *integer) modify() {
	*i++
}

func main() {
	var i integer = 20
	i.print()
	i.modify()
	i.print()

	fmt.Println()

	stu := Student{"Jason", 26}
	//如果實現了*Student類型的String()方法，就會自動調用
	fmt.Println(&stu)
}
