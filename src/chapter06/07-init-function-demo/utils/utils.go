package utils

import "fmt"

var Age int
var Name string

// 初始化變數值給main.go使用
func init() {
	fmt.Println("utils init()...")
	Age = 30
	Name = "Tom"
}
