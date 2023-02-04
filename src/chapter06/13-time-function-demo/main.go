package main

import (
	"fmt"
	"time"
)

// golang中的time
func main() {
	//1. 獲取當前時間
	now := time.Now()
	fmt.Printf("now=%v, type=%T \n", now, now)

	//2. 通過now可以獲得年月日時分秒
	fmt.Printf("年=%v \n", now.Year())
	fmt.Printf("月=%v \n", int(now.Month()))
	fmt.Printf("日=%v \n", now.Day())
	fmt.Printf("時=%v \n", now.Hour())
	fmt.Printf("分=%v \n", now.Minute())
	fmt.Printf("秒=%v \n", now.Second())

	//格式化日期時間
	//1)使用`Printf`或者`SPrintf`
	fmt.Printf("當前時間：%d-%d-%d %d:%d:%d \n", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
	dateStr := fmt.Sprintf("當前時間：%d-%d-%d %d:%d:%d \n", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println("dateStr=", dateStr)

	//2)使用`.Format()`
	fmt.Println(now.Format("2006/01/02 15:04:06"))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("15:04:06"))
	fmt.Println("年：", now.Format("2006"))

	//每隔1秒鐘打印一個數字
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("i=", i)
	}
	//每隔0.1秒鐘打印一個數字
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second / 10)
		//time.Sleep(time.Millisecond * 100)
		fmt.Println("i=", i)
	}

	//獲取當前unix時間戳和unixnano時間戳
	fmt.Println("unix", now.Unix())
	fmt.Println("unixnano", now.UnixNano())

}
