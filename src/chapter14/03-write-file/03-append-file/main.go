package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//打開一個存在的文件，在原先的內容追加1句"ABC!ENGLISH"
	file, err := os.OpenFile("/Users/jason/Desktop/test1.txt", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file error=", err)
	} else {
		defer file.Close()

		writer := bufio.NewWriter(file)
		writer.WriteString("ABC!ENGLISH\n")
		writer.Flush()
	}
}
