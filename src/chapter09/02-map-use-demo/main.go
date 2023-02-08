package main

import "fmt"

func main() {
	//第一種使用方式
	var map1 map[string]string
	map1 = make(map[string]string, 10)
	map1["k1"] = "v1"
	fmt.Println("map1=", map1)

	//第二種使用方式
	var cities = make(map[string]string)
	cities["no1"] = "台北"
	cities["no2"] = "台中"
	cities["no3"] = "台南"
	fmt.Println("cities=", cities)

	//第三種使用方式
	var heroes = map[string]string{
		"no1": "hero1",
		"no2": "hero2",
		"no3": "hero3",
	}
	fmt.Println("heroes=", heroes)
}
