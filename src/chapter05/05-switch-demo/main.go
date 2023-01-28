package main

import "fmt"

func main() {
	var day byte
	fmt.Print("請輸入一個字符a,b,c,d,e,f,g：")
	fmt.Scanf("%c", &day)

	//switch test(day) { // fr
	switch day {
	case 'a':
		fmt.Println("星期一")
	case 'b':
		fmt.Println("星期二")
	case 'c':
		fmt.Println("星期三")
	case 'd':
		fmt.Println("星期四")
	case 'e':
		fmt.Println("星期五")
	case 'f':
		fmt.Println("星期六")
	case 'g':
		fmt.Println("星期日")
	default:
		fmt.Println("輸出有誤")
	}

	//報錯，因為`n1`和`n2`數據類型不一致
	//var n1 int32 = 20
	//var n2 int64 = 20
	//switch n1 {
	//case n2:
	//	fmt.Println("ok1")
	//default:
	//	fmt.Println("沒有匹配")
	//}

	var i1 int32 = 20
	var i2 int32 = 15
	switch i1 {
	case i2, 20: //case後面可以有多個表達式
		fmt.Println("ok!")
	default:
		fmt.Println("沒有匹配")
	}
}

func test(b byte) byte {
	return b + 1
}
