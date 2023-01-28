package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64 = 3.0, 100.0, 6.0

	m := b*b - 4*a*c
	if m > 0 {
		x1 := (-b + math.Sqrt(m)) / 2 * a
		x2 := (-b - math.Sqrt(m)) / 2 * a
		fmt.Printf("x1=%v, x2=%v", x1, x2)
	} else if m == 0 {
		x1 := (-b + math.Sqrt(m)) / 2 * a
		fmt.Printf("x1=%v", x1)
	} else {
		fmt.Println("無解...")
	}
}
