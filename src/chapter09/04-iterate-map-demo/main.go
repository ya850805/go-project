package main

import "fmt"

func main() {
	//使用for-range遍歷map
	var cities = make(map[string]string)
	cities["no1"] = "台北"
	cities["no2"] = "台中"
	cities["no3"] = "台南"

	for k, v := range cities {
		fmt.Printf("key=%v, value=%v \n", k, v)
	}
	fmt.Println("cities有", len(cities), "對key-value")

	fmt.Println()

	//遍歷結構較為複雜的map
	var students = make(map[string]map[string]string)

	students["001"] = make(map[string]string)
	students["001"]["name"] = "Jason"
	students["001"]["sex"] = "M"

	students["002"] = make(map[string]string)
	students["002"]["name"] = "Juan"
	students["002"]["sex"] = "F"

	students["003"] = make(map[string]string)
	students["003"]["name"] = "Marry"
	students["003"]["sex"] = "F"

	for k1, v1 := range students {
		fmt.Println("k1=", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t k2=%v, v2=%v \n", k2, v2)
		}
	}
}
