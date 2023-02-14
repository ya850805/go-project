package main

import "fmt"

type Stu struct {
	Name string
	Age  int
}

func main() {
	//方式1
	//在創建結構體變數時，就直接指定字段的值
	var stu1 = Stu{"小明", 18}
	stu2 := Stu{"小花", 19}

	//在創建結構體變數時，把字段名和字段寫在一起，這種寫法就不依賴字段的定義順序
	var stu3 = Stu{
		Name: "jack",
		Age:  20,
	}
	var stu4 = Stu{
		Age:  30,
		Name: "marry",
	}

	fmt.Println(stu1, stu2, stu3, stu4)

	//方式2：返回結構體變數的指針類型
	var stu5 = &Stu{"小王", 29} // stu5 --> 地址 --> 結構體數據[xxx,xxx]
	stu6 := &Stu{"小陳", 21}

	//在創建結構體指針變數時，把字段名和字段寫在一起，這種寫法就不依賴字段的定義順序
	var stu7 = &Stu{
		Name: "小李",
		Age:  29,
	}
	stu8 := &Stu{
		Age:  31,
		Name: "小張",
	}

	fmt.Println(*stu5, *stu6, *stu7, *stu8)
}
