package main

import "fmt"

func main() {
	var cities = make(map[string]string)
	cities["no1"] = "台北"
	cities["no2"] = "台中"
	cities["no3"] = "台南"
	fmt.Println("cities=", cities)

	//因為no3這個key已經存在，因此下面這句話就是修改
	cities["no3"] = "嘉義"
	fmt.Println("cities=", cities)

	//刪除
	delete(cities, "no1")
	fmt.Println("cities=", cities)
	//如果key不存在，不操作但是也不會報錯
	delete(cities, "no4")

	//map查找
	val, findRes := cities["no3"]
	if findRes {
		fmt.Println("有這個key，值為：", val)
	} else {
		fmt.Println("沒有這個key")
	}

	//如果希望一次性刪除map所有key
	//1.遍歷一下key，逐個刪除
	//2.make一個新的
	cities = make(map[string]string)
	fmt.Println("cities=", cities)

}
