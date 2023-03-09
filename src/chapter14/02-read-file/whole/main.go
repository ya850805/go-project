package main

import (
	"fmt"
	"os"
)

func main() {
	file := "/Users/jason/Desktop/test.txt"

	content, err := os.ReadFile(file)

	if err != nil {
		fmt.Println("read file error=", err)
	} else {
		//需要轉換成string
		fmt.Printf("%v", string(content))

		//我們沒有顯式的Open文件，因此也不需要顯式的Close文件
		//因為，文件的Open和Close也被封裝到ReadFile函數內部
	}
}
