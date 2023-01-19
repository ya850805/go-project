package main

import "fmt"

// golang中的進制
func main() {
	var i int = 5
	//二進制輸出
	fmt.Printf("%b \n", i)

	//**八進制**：0~7，滿8進1，以數字0開頭
	var j int = 011
	fmt.Println("j=", j)

	//**十六進制**：0~9及A~F，滿16進1，以0x或0X開頭表示，此處的A~F不區分大小寫。
	var k int = 0x11
	fmt.Println("k=", k)
}
