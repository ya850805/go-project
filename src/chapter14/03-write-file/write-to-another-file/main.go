package main

import (
	"fmt"
	"os"
)

func main() {
	//編寫一個程序，將一個文件的內容，寫到另外一個文件。注意：這兩個文件已經存在了。
	//使用`os.ReadFile`, `os.WriteFile`完成。
	path1 := "/Users/jason/Desktop/test1.txt"
	path2 := "/Users/jason/Desktop/test2.txt"

	//讀取
	content, err := os.ReadFile(path1)
	if err != nil {
		fmt.Println("read file error=", err)
	} else {
		//寫入
		err = os.WriteFile(path2, content, 0666)
		if err != nil {
			fmt.Println("write file error=", err)
		}
	}
}
