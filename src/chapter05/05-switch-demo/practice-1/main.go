package main

import "fmt"

func main() {
	//使用`switch`把小寫類型的char型轉為大寫(鍵盤輸入)。只轉換a,b,c,d,e，其他的輸出other。
	var letter byte
	fmt.Print("請輸入字母a,b,c,d,e：")
	fmt.Scanf("%c", &letter)

	switch letter {
	case 'a':
		fmt.Println("A")
	case 'b':
		fmt.Println("B")
	case 'c':
		fmt.Println("C")
	case 'd':
		fmt.Println("D")
	case 'e':
		fmt.Println("E")
	default:
		fmt.Println("other")
	}
}
