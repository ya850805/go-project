package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//打開一個存在的文件，將原來的內容覆蓋成新的內容10句"hello jason"
	file, err := os.OpenFile("/Users/jason/Desktop/test1.txt", os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("open file error=", err)
	} else {
		defer file.Close()

		writer := bufio.NewWriter(file)
		for i := 0; i < 10; i++ {
			writer.WriteString("hello jason\n")
		}

		writer.Flush()
	}
}
