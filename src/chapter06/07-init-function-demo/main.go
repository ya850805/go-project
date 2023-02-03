package main

import (
	"fmt"
	"go-project/src/chapter06/07-init-function-demo/utils"
)

var age = test()

// 為了看到全局變數是先被初始化的，我們寫一個函數
func test() int {
	fmt.Println("test()...")
	return 90
}

// init函數，通常可以在init函數中完成初始化工作
func init() {
	fmt.Println("init()...")
}

func main() {
	fmt.Println("main()...")
	fmt.Println("utils Age=", utils.Age)
	fmt.Println("utils Name=", utils.Name)
}
