package main

import "fmt" //fmt包中提供格式化、輸出、輸入的函數

func main() {
	fmt.Println("jason\tjuan")

	fmt.Println("hello\nworld")

	fmt.Println("\\")

	fmt.Println("jason:\"hello\"")

	// \r回車，從當前行的最前面開始輸出，覆蓋掉以前的內容
	fmt.Println("hello\rworld!")
	fmt.Println("second\rfirst")

	fmt.Println("姓名\t年齡\t籍貫\t住址\njohn\t12\t台灣\t台北")
}
