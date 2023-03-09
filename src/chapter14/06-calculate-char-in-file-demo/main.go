package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	ChCount    int //英文
	NumCount   int //數字
	SpaceCount int //空格
	OtherCount int //其他
}

func main() {
	//打開一個文件，創建一個Reader
	//每讀取一行，就去統計該行有多少個英文、數字、空格和其他字符
	//然後將結果保存到一個結構體

	file, err := os.Open("/Users/jason/Desktop/abc.txt")
	if err != nil {
		fmt.Println("open file error=", err)
		return
	}
	defer file.Close()

	//結構體
	charCount := CharCount{}
	reader := bufio.NewReader(file)
	for {
		str, error := reader.ReadString('\n')
		if error == io.EOF {
			break
		}

		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				charCount.ChCount++
			case v == ' ' || v == '\t':
				charCount.SpaceCount++
			case v >= '0' && v <= '9':
				charCount.NumCount++
			default:
				charCount.OtherCount++
			}
		}
	}

	fmt.Println(charCount)
}
