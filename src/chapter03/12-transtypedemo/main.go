package main

import "fmt"

// golang中數據類型的轉換
func main() {
	var i int32 = 100

	//將i轉成float32
	var n1 float32 = float32(i)

	var n2 int8 = int8(n1)

	var n3 int64 = int64(i) //低精度->高精度，還是要強制轉換

	fmt.Printf("i=%d, n1=%f, n2=%d, n3=%d \n", i, n1, n2, n3)

	//被轉換的是變數存儲的數據(即值)，變數本身的數據類型並沒有變化
	fmt.Printf("i的類型是：%T \n", i)

	//在轉換中，例如將int64轉成int8(-128~127)，編譯時不會報錯，
	//只是轉換的結果是按溢出處理，和我們希望的結果不一樣
	var num1 int64 = 999999
	var num2 int8 = int8(num1)
	fmt.Println("num2=", num2)
}
