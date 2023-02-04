package main

import (
	"fmt"
	"strconv"
	"time"
)

func test03() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}

func main() {
	//編寫一段代碼來統計函數`test03`執行的時間
	time1 := time.Now().Unix()
	test03()
	time2 := time.Now().Unix()
	fmt.Printf("test03()總共執行%d秒", time2-time1)
}
