package main

import "fmt"

func main() {
	//打印10句`Hello world!`
	for i := 1; i <= 10; i++ {
		fmt.Println("Hello world!", i)
	}

	//for循環的第二種寫法：將變數初始化和變數迭代寫到其他位置
	j := 1        //循環變數初始化
	for j <= 10 { // 循環條件
		fmt.Println("Hi", j) //循環體
		j++                  //循環變數迭代
	}

	//for循環的第三種寫法：同長配合break使用
	k := 1
	for {
		if k <= 10 {
			fmt.Println("ok~")
		} else {
			break //跳出整個for循環
		}
		k++
	}
}
