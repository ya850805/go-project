package main

import (
	"fmt"
	"reflect"
)

// Monster 結構體
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

func (m Monster) Print() {
	fmt.Println("-----start-----")
	fmt.Println(m)
	fmt.Println("-----end-----")
}

func (m Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (m Monster) Set(name string, age int, score float32, sex string) {
	m.Name = name
	m.Age = age
	m.Score = score
	m.Sex = sex
}

func TestStruct(b interface{}) {
	typ := reflect.TypeOf(b)
	val := reflect.ValueOf(b)
	kd := val.Kind()

	//如果傳入的不是結構體，就退出
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	//獲取到該結構體有幾個字段
	num := val.NumField()
	fmt.Printf("struct has %d fields \n", num) //4

	//遍歷結構體的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d，值為=%v \n", i, val.Field(i))
		//獲取結構體的標籤(透過reflect.Type)
		tagVal := typ.Field(i).Tag.Get("json")
		//如果該字段有標籤tag就顯示，否則就不顯示
		if tagVal != "" {
			fmt.Printf("Field %d，tag為=%v \n", i, tagVal)
		}
	}

	//獲取到該結構體有幾個方法
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods \n", numOfMethod)

	//獲取到第二個方法，並調用他
	//方法的排序默認是按照 函數名的排序(ASCII code)
	//因此這個會調用到Print()
	val.Method(1).Call(nil)

	//調用結構體的第一個方法(GetSum)
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params)
	fmt.Println("res=", res[0].Int())
}

func main() {
	var m = Monster{
		Name:  "monster1",
		Age:   400,
		Score: 30.8,
	}
	TestStruct(m)
}
