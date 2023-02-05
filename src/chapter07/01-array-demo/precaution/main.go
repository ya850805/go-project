package main

import "fmt"

func test01(array [3]int) {
	array[0] = 88
}

func test02(array *[3]int) {
	(*array)[0] = 99
}

func main() {
	//陣列是多個**相同類型數據**的組合，一個陣列一但聲明/定義了，其**長度是固定的，不能動態變化**
	var arr01 [3]int
	arr01[0] = 1
	arr01[1] = 30
	//arr01[2] = 1.1 //這裡會報錯，類型不符
	arr01[2] = 5
	//arr01[3] = 90 //這裡會報錯，因為長度是固定的，不能動態變化，否則報越界

	//陣列創建後，如果沒有賦值，元素會有默認值。
	//* 數值類型陣列：默認值為`0`。
	//* 字符串陣列：默認值為`""`。
	//* `bool`陣列：默認值為`false`。
	var floatArr [3]float32
	var stringArr [3]string
	var boolArr [3]bool
	fmt.Printf("floatArr=%v, stringArr=%v, boolArr=%v \n", floatArr, stringArr, boolArr)

	//Go的陣列屬於值類型，在默認情況下是值傳遞，因此會進行值拷貝，陣列間不會相互影響
	array := [3]int{11, 22, 33}
	test01(array)
	fmt.Println("array=", array)
	test02(&array)
	fmt.Println("array=", array)
}
