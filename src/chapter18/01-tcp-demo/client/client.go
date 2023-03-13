package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}

	for {
		//客戶端發送單行數據，然後就退出
		fmt.Print("請輸入：")
		reader := bufio.NewReader(os.Stdin) //os.Stdin代表標準輸入(終端)

		//從終端讀取一行用戶輸入，並準備發送給服務器
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err=", err)
		}

		str = strings.Trim(str, "\r\n")
		if str == "exit" {
			fmt.Println("客戶端退出...")
			return
		}

		//再將str發送給服務器
		n, err := conn.Write([]byte(str))
		if err != nil {
			fmt.Println("conn Write err=", err)
		}
		fmt.Printf("客戶端發送了%d字節的數據\n", n)
	}
}
