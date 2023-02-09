package main

import "fmt"

func main() {
	//使用一個`map`來記錄monster的信息name和`age`，也就是說一個monster對應一個`map`，並且monster的個數可以動態的增加:point_right:`map`切片。
	var slice []map[string]string
	slice = make([]map[string]string, 2)

	monster1 := make(map[string]string)
	monster1["name"] = "monster A"
	monster1["age"] = "100"
	slice[0] = monster1

	monster2 := make(map[string]string)
	monster2["name"] = "monster B"
	monster2["age"] = "200"
	slice[1] = monster2

	fmt.Println("slice=", slice)

	//如果還需要添加一個monster，需要使用append()，因為一開始定義切片的長度為2
	monster3 := make(map[string]string)
	monster3["name"] = "monster C"
	monster3["age"] = "600"
	slice = append(slice, monster3)

	fmt.Println("slice=", slice)
}
