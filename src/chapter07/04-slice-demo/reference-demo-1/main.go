package main

import "fmt"

func main() {
	//切片是引用類型，所以在傳遞時，遵守引用傳遞機制
	var slice []int
	var arr = [5]int{1, 2, 3, 4, 5}
	slice = arr[:]

	slice2 := slice
	slice2[0] = 10

	fmt.Println("arr=", arr)
	fmt.Println("slice=", slice)
	fmt.Println("slice2=", slice2)
}
