package main

import (
	"fmt"
	"reflect"
)

func reflect01(b interface{}) {
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal的Kind=%v \n", rVal.Kind())
	//修改值需要使用Elem()
	rVal.Elem().SetInt(20)
}

type Student struct {
	Name string
}

func reflect02(b interface{}) {
	rVal := reflect.ValueOf(b)
	rVal.Elem().SetString("Juan")
}

func main() {
	//通過反射修改

	//num int的值
	var num int = 10
	reflect01(&num)
	fmt.Println("num=", num)

	//student的值
	stu := Student{Name: "Jason"}
	reflect02(&(stu.Name))
	fmt.Println("stu.Name=", stu.Name)
}
