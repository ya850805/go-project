package main

import "fmt"

// Student 學生
type Student struct {
	Name  string
	Age   int
	Score int
}

func (stu *Student) showInfo() {
	fmt.Printf("學生名=%v 年齡=%v 成績=%v \n", stu.Name, stu.Age, stu.Score)
}

// Pupils 小學生
type Pupils struct {
	Student //嵌入結構體
}

// Graduate 大學生
type Graduate struct {
	Student //嵌入結構體
}

func (stu *Student) SetScore(score int) {
	stu.Score = score
}

// Pupils特有的方法
func (p *Pupils) testing() {
	fmt.Println("小學生正在考試中...")
}

// Graduate特有的方法
func (p *Graduate) testing() {
	fmt.Println("大學生正在考試中...")
}

func main() {
	//小學生
	pupil := &Pupils{}
	pupil.Student.Name = "小學生1"
	pupil.Student.Age = 10

	pupil.testing()
	pupil.Student.SetScore(70)
	pupil.Student.showInfo()

	//大學生
	graduate := &Graduate{}
	graduate.Student.Name = "大學生1"
	graduate.Student.Age = 20

	graduate.testing()
	graduate.Student.SetScore(90)
	graduate.Student.showInfo()
}
