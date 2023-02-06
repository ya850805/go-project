package main

import "fmt"

func test(slice []int) {
	slice[0] = 100 //因為是引用傳遞，這裡修改slice[0]，會改變實際切片
}

func main() {
	//切片是引用類型，所以在傳遞時，遵守引用傳遞機制
	var slice = []int{1, 2, 3, 4}
	fmt.Println("slice=", slice)
	test(slice)
	fmt.Println("slice=", slice)
}
