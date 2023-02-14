package main

import "fmt"

type Student struct {
	Name   string
	Gender string
	Age    int
	Id     int
	Score  float64
}

func (s *Student) say() string {
	return fmt.Sprintf("信息如下：Name=[%v], Gender=[%v], Age=[%v], Id=[%v], Score=[%v]", s.Name, s.Gender, s.Age, s.Id, s.Score)
}

func main() {
	//1. 編寫一個`Student`結構體，包含`Name`、`Gender`、`Age`、`Id`、`Score`字段，分別為`string`、`string`、`int`、`int`、`float64`類型。
	//2. 結構體中聲明一個`say`方法，返回`string`類型，方法返回信息中包含所有字段值。
	//3. 在`main`方法中，創建`Student`結構體實例(變數)，並訪問`say`方法，並將調用結果打印輸出。
	student := Student{Name: "Jason", Gender: "Male", Age: 26, Id: 1000, Score: 99.98}
	fmt.Println(student.say())
}
