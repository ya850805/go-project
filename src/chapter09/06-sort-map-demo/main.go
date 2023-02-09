package main

import (
	"fmt"
	"sort"
)

func main() {
	//map的排序
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90

	//如何按照map的key順序進行排序輸出
	//1. 先將map的key放入到切片中
	//2. 對切片排序
	//3. 遍歷切片，然後按照key來輸出map的值
	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k)
	}

	//排序
	sort.Ints(keys)
	fmt.Println("keys=", keys)

	for _, v := range keys {
		fmt.Printf("map[%v]=%v \n", v, map1[v])
	}
}
