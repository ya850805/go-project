package main

import "fmt"

// 請編寫一個函數`swap(n1 *int, n2 *int)`可以交換`n1`和`n2`的值。
func swap(n1 *int, n2 *int) {
	n3 := *n1
	*n1 = *n2
	*n2 = n3
}

func main() {
	n1 := 10
	n2 := 20

	swap(&n1, &n2)
	fmt.Printf("n1=%v, n2=%v", n1, n2)
}
