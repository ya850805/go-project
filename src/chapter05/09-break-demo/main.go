package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//隨機生成1~100的一個數，直到生成了99這個數，看看一共生成了幾個數？
	var count = 0

	//生成隨機數需要設置rand.Seed()
	// Seeding with the same value results in the same random sequence each run.
	// For different numbers, seed with a different value, such as
	// time.Now().UnixNano(), which yields a constantly-changing number.
	rand.Seed(time.Now().Unix())
	for {
		//隨機生成1~100的一個數
		var num = rand.Intn(100) + 1
		fmt.Println(num)
		count++
		if num == 99 {
			break
		}
	}

	fmt.Printf("一共生成了%d個數", count)
}
