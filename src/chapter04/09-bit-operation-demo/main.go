package main

import "fmt"

//原碼、反碼、補碼

//對於有符號的數字而言：

//二進制的最高位是符號位，0表示正數，1表示負數。
//1 :point_right: 0000 0001
//-1 :point_right: 1000 0001

//正數的原碼、反碼、補碼都一樣。

//負數的反碼=他的原碼符號位不變，其他位取反(0 -> 1 / 1 -> 0)。
//1:point_right:原碼0000 0001反碼0000 0001補碼0000 0001
//-1:point_right:原碼1000 0001反碼1111 1110補碼1111 1111

//負數的補碼=他的反碼+1。

//0的反碼、補碼都是0。

//在電腦運行的時候，都是以補碼的方式來運算的。
//1 - 1 = 1 + (-1)

func main() {
	//位運算
	fmt.Println(2 & 3)
	fmt.Println(2 | 3)
	fmt.Println(2 ^ 3)

	//負數
	fmt.Println(-2 ^ 2)

	//移位
	fmt.Println(1 >> 2)
	fmt.Println(1 << 2)
}
