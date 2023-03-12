package main

import (
	"fmt"
	"reflect"
)

func test(b interface{}) {
	rVal := reflect.ValueOf(b)

	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)

	rKind := rVal.Kind()
	fmt.Println("rKind=", rKind)

	value1 := rVal.Float()
	fmt.Println("value1=", value1)

	i := rVal.Interface()
	value2 := i.(float64)
	fmt.Println("value2=", value2)
}

func main() {
	//給你一個變數`var v float64 = 1.2`，請使用反射來得到他的`reflect.Value`，
	//然後獲取對應的Type和Kind和值，並將`reflect.Value`轉換程`interface{}`，再將`interface{}`轉換成`float64`

	var v float64 = 1.2
	test(v)
}
