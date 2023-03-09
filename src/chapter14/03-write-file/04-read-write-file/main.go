package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//打開一個存在的文件，把原來的內容**讀出**顯示在終端，並**追加**5句"你好"
	file, err := os.OpenFile("/Users/jason/Desktop/test1.txt", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file error=", err)
	} else {
		defer file.Close()

		//讀取
		reader := bufio.NewReader(file)
		for {
			str, error := reader.ReadString('\n')
			if error == io.EOF { //文件的末尾
				break
			}
			//顯示到終端
			fmt.Print(str)
		}

		//寫入
		writer := bufio.NewWriter(file)
		for i := 0; i < 5; i++ {
			writer.WriteString("你好\n")
		}
		writer.Flush()
	}
}
