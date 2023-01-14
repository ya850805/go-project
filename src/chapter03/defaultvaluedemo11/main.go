package main

import "fmt"

// golang中默認值
func main() {
	var a int
	var b float32
	var c float64
	var isMarried bool
	var name string

	// %v表示直接按照變數的值輸出
	fmt.Printf("a=%d, b=%f, c=%f, isMarried=%v, name=%v", a, b, c, isMarried, name)
}
