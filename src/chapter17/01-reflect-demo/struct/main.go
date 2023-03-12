package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func reflectTest01(b interface{}) {
	//通過反射後取傳入的變數的type, kind, 值

	//1. 獲取reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)

	//2. 獲取reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal=", rVal)
	fmt.Printf("rVal的類型=%T \n", rVal)

	//3. 將reflect.Value轉成interface{}
	iV := rVal.Interface()
	fmt.Printf("iV=%v，iV type=%T \n", iV, iV)
	//將interface{}通過斷言轉成需要的類型(這樣才能取到Student的屬性)
	//這裡就簡單使用了一個帶檢測的類型斷言
	//我們也可以使用switch的斷言形式來做更加靈活的判斷
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v \n", stu.Name)
	}

	//4. 獲取reflect.Kind
	kind1 := rType.Kind()
	kind2 := rVal.Kind()
	fmt.Printf("kind=%v kind=%v", kind1, kind2)
}

func main() {
	//演示對(結構體類型、`interface{}`、`reflect.Value`)進行反射的基本操作

	stu := Student{
		Name: "Jason",
		Age:  26,
	}
	reflectTest01(stu)
}
