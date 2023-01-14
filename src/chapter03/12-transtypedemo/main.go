package main

import "fmt"

// golang中數據類型的轉換
func main() {
	var i int32 = 100

	//將i轉成float32
	var n1 float32 = float32(i)

	var n2 int8 = int8(n1)

	var n3 int64 = int64(i) //低精度->高精度，還是要強制轉換

	fmt.Printf("i=%d, n1=%f, n2=%d, n3=%d", i, n1, n2, n3)
}
