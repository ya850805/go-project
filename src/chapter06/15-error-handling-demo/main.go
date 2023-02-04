package main

import (
	"errors"
	"fmt"
)

//golang的錯誤處理

func test() {
	//使用`defer`+`recover`來捕獲、處理錯誤
	defer func() {
		err := recover() //recover()內置函數，可以捕獲到異常
		if err != nil {  //捕獲到錯誤
			//這裡就做異常處理
			fmt.Println(err)
		}
	}()

	num1 := 10
	num2 := 0

	res := num1 / num2
	fmt.Println("res=", res)
}

// 模擬函數去讀取一個配置文件的信息
// 如果文件名傳入不正確，就拋出一個自定義錯誤
func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		//返回自定義錯誤
		return errors.New("讀取文件錯誤")
	}
}

func test02() {
	err := readConf("configXXX.ini")
	if err != nil {
		//如果讀取文件函數返回錯誤，就輸出這個錯誤，並終止程序
		panic(err)
	}
	fmt.Println("test02()繼續執行...")
}

func main() {
	test()
	fmt.Println("main下方的代碼")

	//測試自定義錯誤
	test02()
}
