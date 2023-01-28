package main

import "fmt"

func main() {
	var height int32
	var money float32
	var isHandsome bool

	fmt.Print("請輸入身高(公分)：")
	fmt.Scanln(&height)

	fmt.Print("請輸入財富(千萬)：")
	fmt.Scanln(&money)

	fmt.Print("是否長得帥(true/false)：")
	fmt.Scanln(&isHandsome)

	if height > 180 && money > 1.0 && isHandsome {
		fmt.Println("我一定要嫁給他")
	} else if height > 180 || money > 1.0 || isHandsome {
		fmt.Println("嫁吧，比上不足，比下有餘")
	} else {
		fmt.Println("不嫁")
	}
}
