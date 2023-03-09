package main

import (
	"fmt"
	"os"
)

func main() {
	//打開一個文件
	file, err := os.Open("/Users/jason/Desktop/test.txt")
	if err != nil {
		fmt.Println("open file error=", err)
	} else {
		//輸出看文件是什麼，看出file就是一個指針*File
		fmt.Printf("file=%v \n", file)

		err = file.Close()
		if err != nil {
			fmt.Println("close file error=", err)
		}
	}
}
