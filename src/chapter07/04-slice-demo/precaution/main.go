package main

import "fmt"

func main() {
	var arr = [...]int{10, 20, 30, 40, 50}
	slice := arr[1:4]

	//切片可以再切片
	slice2 := slice[1:2]
	fmt.Println("slice=", slice)
	fmt.Println("slice2=", slice2)

	fmt.Println()

	//因為arr, slice, slice2指向的數據空間是同一個，因此會一次全部修改
	slice2[0] = 100
	fmt.Println("arr=", arr)
	fmt.Println("slice=", slice)
	fmt.Println("slice2=", slice2)

	fmt.Println()

	//用`append()`內置函數可以對切片進行動態追加
	var slice3 = []int{100, 200, 300}

	//通過append給slice3追加具體的元素
	slice3 = append(slice3, 400, 500, 600)
	fmt.Println("slice3=", slice3)

	//通過append將切片slice3(可以是任意切片)追加給slice3
	slice3 = append(slice3, slice3...)
	fmt.Println("slice3=", slice3)

	fmt.Println()

	//切片的拷貝操作
	var slice4 = []int{1, 2, 3, 4, 5}
	var slice5 = make([]int, 10)
	copy(slice5, slice4)
	fmt.Println("slice4=", slice4)
	fmt.Println("slice5=", slice5)
}
