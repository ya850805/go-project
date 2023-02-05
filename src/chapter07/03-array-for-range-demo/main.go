package main

import "fmt"

// for-range遍歷陣列
func main() {
	var heroes = [3]string{"hero A", "hero B", "hero C"}

	for index, hero := range heroes {
		fmt.Printf("下標=%v, 值=%v \n", index, hero)
	}
}
