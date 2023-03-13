package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()

	//循環接收客戶端發送的數據
	for {
		//創建一個新的切片
		buffer := make([]byte, 1024)

		//conn.Read(buffer)
		//1. 等待客戶端通過conn發送信息
		//2. 如果客戶端沒有write(發送)，那麼協程就會阻塞在這裏
		//fmt.Printf("服務器在等待客戶端%s發送信息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buffer) //從conn讀取
		if err != nil {
			fmt.Printf("server Read() err=%v \n", err)
			return
		}

		//顯示客戶端發送的數據在服務器終端
		//需要指定長度
		fmt.Println(string(buffer[:n]))
	}
}

func main() {
	fmt.Println("服務器開始監聽...")
	//tcp表示網路協議是tcp
	//0.0.0.0:8888表示在本地監聽8888端口
	listen, err := net.Listen("tcp", "0.0.0.0:8888")

	if err != nil {
		fmt.Println("listen err=", err)
		return
	}

	//延時關閉listen
	defer listen.Close()

	for {
		//等待客戶端連接
		fmt.Println("等待客戶端來連接...")

		//這邊會卡著，等待連接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
			return
		} else {
			fmt.Printf("Accept() success，客戶端IP=%v \n", conn.RemoteAddr().String())
			//創建一個協程，為客戶端服務
			go process(conn)
		}

	}
}
