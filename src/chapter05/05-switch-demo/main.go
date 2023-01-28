package main

import "fmt"

func main() {
	var day byte
	fmt.Print("請輸入一個字符a,b,c,d,e,f,g：")
	fmt.Scanf("%c", &day)

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
}
