package main

import "fmt"

func main() {
	//map的聲明和注意事項
	var a map[string]string
	//在使用map前，需要先make，make的作用就是給map分配數據空間
	a = make(map[string]string, 10)
	a["k1"] = "v1"
	a["k2"] = "v2"
	a["k1"] = "value1"
	a["k3"] = "v2"
	fmt.Println(a)
}
