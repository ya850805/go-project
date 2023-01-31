package main

import "fmt"

func main() {
	//指定標前來使用break

outer:
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				//break默認跳出最近的for循環
				//指定標籤可以指定break要跳出的for循環
				break outer
			}
			fmt.Println("j=", j)
		}
	}
}
