package main

import "fmt"

func main() {
	//`switch`後也可以不帶表達式，類似`if-else`分支來使用
	var age int = 10
	switch {
	case age == 10:
		fmt.Println("age=10")
	case age == 20:
		fmt.Println("age=20")
	}

	//也可以對範圍進行判斷
	var score int = 95
	switch {
	case score > 90:
		fmt.Println("成績優良..")
	case score >= 60 && score <= 90:
		fmt.Println("成績及格..")
	default:
		fmt.Println("成績不及格..")
	}

	//`switch`後也可以直接聲明/定義一個變數，分號結束，不推薦
	switch grade := 90; {
	case grade > 90:
		fmt.Println("成績優良..")
	case grade >= 60 && grade <= 90:
		fmt.Println("成績及格..")
	default:
		fmt.Println("成績不及格..")
	}

	//`switch`穿透`fallthrough`，如果在`case`語句後增加`fallthrough`，則**會繼續執行下一個`case`，也叫`switch`穿透**
	var num int = 10
	switch num {
	case 10:
		fmt.Println("ok1")
		fallthrough //默認只能穿透一層(因此下面case2不會做判斷，ok2也會直接打印)
	case 20:
		fmt.Println("ok2")
	case 30:
		fmt.Println("ok3")
	default:
		fmt.Println("沒有匹配")
	}

	//**Type switch**：`switch`語句還可以被用於type-switch來判斷某個interface變數中實際指向的變數類型。
	var x interface{}
	var y = 10.0
	x = y
	switch x.(type) {
	case nil:
		fmt.Println("nil")
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case bool, string:
		fmt.Println("bool, string")
	default:
		fmt.Println("unknown..")
	}
}
