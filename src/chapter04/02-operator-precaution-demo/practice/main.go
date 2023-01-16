package main

import "fmt"

func main() {
	//假如還有97天放假，問：還有xx個星期xx天？
	var days int = 97
	fmt.Println("還有", days/7, "個星期，", days%7, "天")

	//定義一個變數保存華氏溫度，華氏溫度轉換攝氏溫度的公式為：`5 / 9 * (華氏溫度 - 100)`，請求出華氏溫度對應的攝氏溫度。
	var temp float32 = 134.2
	fmt.Println("攝氏溫度：", 5.0/9*(temp-100))
}
