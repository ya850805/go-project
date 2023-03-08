package main

import "fmt"

type Student struct {
}

// TypeJudge 編寫一個函數，可以判斷傳入的參數是什麼類型
func TypeJudge(items ...interface{}) {
	for index, x := range items {
		index++
		switch x.(type) {
		case bool:
			fmt.Printf("第%v個參數是bool類型，值是%v \n", index, x)
		case float32:
			fmt.Printf("第%v個參數是float32類型，值是%v \n", index, x)
		case float64:
			fmt.Printf("第%v個參數是float64類型，值是%v \n", index, x)
		case int, int32, int64:
			fmt.Printf("第%v個參數是整數類型，值是%v \n", index, x)
		case string:
			fmt.Printf("第%v個參數是string類型，值是%v \n", index, x)
		case Student:
			fmt.Printf("第%v個參數是Student類型，值是%v \n", index, x)
		case *Student:
			fmt.Printf("第%v個參數是Student指針類型，值是%v \n", index, x)
		default:
			fmt.Printf("第%v個參數類型不確定，值是%v \n", index, x)
		}
	}
}

func main() {
	var n1 float32 = 1.1
	var n2 float64 = 2.3
	var n3 int32 = 30
	var name string = "tom"
	address := "TPE"
	n4 := 300

	stu1 := Student{}
	stu2 := &Student{}

	TypeJudge(n1, n2, n3, name, address, n4, stu1, stu2)
}
