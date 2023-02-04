package main

import "fmt"

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

func main() {
	test()
	fmt.Println("main下方的代碼")
}
