package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// CopyFile 編寫一個函數，接收兩個文件路徑srcFileName dstFileName
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	//Reader
	srcFile, err1 := os.Open(srcFileName)
	if err1 != nil {
		fmt.Println("open source file error=", err1)
		return
	}
	reader := bufio.NewReader(srcFile)

	//Writer
	dstFile, err2 := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err2 != nil {
		fmt.Println("open dest file error=", err2)
		return
	}
	writer := bufio.NewWriter(dstFile)

	defer srcFile.Close()
	defer dstFile.Close()

	//Copy
	return io.Copy(writer, reader)
}

func main() {
	//拷貝圖片文件
	_, err := CopyFile("/Users/jason/Desktop/leetcode2.png", "/Users/jason/Desktop/leetcode.png")
	if err == nil {
		fmt.Println("拷貝完成")
	} else {
		fmt.Println("拷貝失敗：", err)
	}
}
