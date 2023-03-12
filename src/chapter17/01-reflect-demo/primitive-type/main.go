package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	//通過反射後取傳入的變數的type, kind, 值

	//1. 獲取reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)

	//2. 獲取reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal=", rVal)
	fmt.Printf("rVal的類型=%T \n", rVal)

	//取值並計算
	n1 := 2 + rVal.Int()
	fmt.Println("n1=", n1)

	//3. 將reflect.Value轉成interface{}
	iV := rVal.Interface()
	//將interface{}通過斷言轉成需要的類型
	num := iV.(int)
	fmt.Println("num=", num)
}

func main() {
	//演示對(基本數據類型、`interface{}`、`reflect.Value`)進行反射的基本操作

	//1. 先定義一個int
	var num int = 100
	reflectTest01(num)
}
