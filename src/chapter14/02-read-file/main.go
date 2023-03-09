package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("/Users/jason/Desktop/test.txt")

	if err != nil {
		fmt.Println("open file error=", err)
	} else {
		//當函數退出時，要及時的關閉file
		//否則會有內存泄露
		defer file.Close()

		//創建一個*Reader，是帶緩衝的
		//默認緩衝區為4096個字節
		reader := bufio.NewReader(file)

		for {
			str, err := reader.ReadString('\n') //讀到換一行就結束
			if err == io.EOF {                  //EOF表示文件的末尾
				break
			}

			//輸出內容
			fmt.Print(str)
		}

		fmt.Println("文件讀取結束...")
	}
}
