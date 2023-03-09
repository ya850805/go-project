package main

import (
	"flag"
	"fmt"
)

func main() {
	//定義幾個變數，用於接收命令行的參數值
	var user string
	var pwd string
	var host string
	var port int

	//&user 就是接收用戶命令行中的-u後面的參數值
	//"u" 就是-u指定參數
	//"" 默認值
	//"用戶名，默認為空" 參數的說明
	flag.StringVar(&user, "u", "", "用戶名，默認為空")
	flag.StringVar(&pwd, "p", "", "密碼名，默認為空")
	flag.StringVar(&host, "h", "localhost", "主機名，默認為localhost")
	flag.IntVar(&port, "P", 3306, "端口好，默認為3306")

	//轉換，必須要用該方法
	flag.Parse()

	fmt.Printf("user=%v, pwd=%v, host=%v, port=%v", user, pwd, host, port)
}
