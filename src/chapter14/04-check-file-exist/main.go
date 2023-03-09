package main

import (
	"fmt"
	"os"
)

func PathExist(path string) (bool, error) {
	_, err := os.Stat(path)

	//文件存在
	if err == nil {
		return true, nil
	}

	//文件不存在
	if os.IsNotExist(err) {
		return false, nil
	}

	//其他錯誤
	return false, err
}

func main() {
	file1Exist, _ := PathExist("/Users/jason/Desktop/test1.txt")
	file2Exist, _ := PathExist("/Users/jason/Desktop/test3.txt")

	fmt.Println(file1Exist)
	fmt.Println(file2Exist)
}
