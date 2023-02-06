package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//隨機生成五個數，並將其反轉打印
	var array [5]int
	len := len(array)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		array[i] = rand.Intn(100) // 0 <= n < 100
	}

	fmt.Println("反轉前array=", array)

	//反轉打印，交換的次數是len / 2，第一個和倒數第一個元素交換，第二個和倒數第二個元素交換...
	for i := 0; i < len/2; i++ {
		temp := array[len-1-i]
		array[len-1-i] = array[i]
		array[i] = temp
	}

	fmt.Println("反轉後array=", array)
}
