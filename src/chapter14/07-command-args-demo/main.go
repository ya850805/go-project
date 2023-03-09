package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("命令行的參數有", len(os.Args), "個")

	//遍歷os.Args切片，就可以得到所有命令行參數值
	for i, arg := range os.Args {
		fmt.Printf("args[%v]=%v \n", i, arg)
	}
}
