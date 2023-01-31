package main

import "fmt"

// golang中的goto使用案例
func main() {
	fmt.Println("ok1")
	goto abc //goto通常搭配if語句使用
	fmt.Println("ok2")
	fmt.Println("ok3")
	fmt.Println("ok4")
abc:
	fmt.Println("ok5")
	fmt.Println("ok6")
	fmt.Println("ok7")
}
