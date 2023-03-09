package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//創建一個新文件，寫入內容5句"hello world"
	file, err := os.OpenFile("/Users/jason/Desktop/test1.txt", os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println("open file error=", err)
	} else {
		//關閉file
		defer file.Close()

		str := "hello world\n"
		//寫入時，使用帶緩衝的*Writer
		writer := bufio.NewWriter(file)
		for i := 0; i < 5; i++ {
			writer.WriteString(str)
		}

		//因為writer是帶緩衝的，因此在調用WriteString內容是先寫到緩衝區的
		//所以需要調用Flush方法，將緩衝的數據真正寫到文件中，否則文件中不會有數據
		writer.Flush()
	}
}
