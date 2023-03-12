package main

import (
	"fmt"
	"reflect"
)

type Cal struct {
	Num1 int
	Num2 int
}

func (c Cal) GetSub(name string) {
	fmt.Printf("%s完成了減法，%d-%d=%d \n", name, c.Num1, c.Num2, c.Num1-c.Num2)
}

func main() {
	//1. 編寫一個`Cal`結構體，有兩個字段`Num1`、`Num2`。
	//2. 方法`GetSub(name string)`。
	//3. 使用反射遍歷`Cal`結構體所有的字段信息。
	//4. 使用反射機制完成對`GetSub`的調用，輸出形式為"tom完成了減法，8-3=5"。
	cal := Cal{
		Num1: 8,
		Num2: 3,
	}

	val := reflect.ValueOf(cal)
	typ := reflect.TypeOf(cal)
	numField := val.NumField()

	//iterate field
	for i := 0; i < numField; i++ {
		fmt.Printf("第%d個字段名稱為：%s，值為：%v \n", i, typ.Field(i).Name, val.Field(i))
	}

	var params []reflect.Value
	params = append(params, reflect.ValueOf("tom"))
	val.Method(0).Call(params)
}
