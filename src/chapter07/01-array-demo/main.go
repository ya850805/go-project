package main

import "fmt"

// golang中的陣列
func main() {
	//一個養雞場有6隻雞，他們的體重分別是3kg, 5kg, 1kg, 3.4kg, 2kg, 50kg。請問這6隻雞的總體重是多少？平均體重是多少？

	//傳統方法
	hen1 := 3.0
	hen2 := 5.0
	hen3 := 1.0
	hen4 := 3.4
	hen5 := 2.0
	hen6 := 50.0

	totalWeight := hen1 + hen2 + hen3 + hen4 + hen5 + hen6
	//將結果四捨五入，保留小數點兩位，並返回
	avgWeight := fmt.Sprintf("%.2f", totalWeight/6)
	fmt.Printf("totalWeight=%v, avgWeight=%v \n", totalWeight, avgWeight)

	//使用陣列的方式
	//1. 定義一個陣列
	var hens [6]float64
	//2. 給陣列的每個元素賦值，下標從0開始
	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 50.0
	//3. 遍歷陣列求出總體重
	totalWeight2 := 0.0
	for i := 0; i < len(hens); i++ {
		totalWeight2 += hens[i]
	}
	//4. 求出平均體重
	avgWeight2 := fmt.Sprintf("%.2f", totalWeight2/float64(len(hens)))
	fmt.Printf("totalWeight2=%v, avgWeight2=%v \n", totalWeight2, avgWeight2)

	//四種初始化陣列的方式
	//1.
	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Println("numArr01=", numArr01)
	//2.
	var numArr02 = [3]int{5, 6, 7}
	fmt.Println("numArr02=", numArr02)
	//3. [...]是固定寫法
	var numArr03 = [...]int{8, 9, 10}
	fmt.Println("numArr03=", numArr03)
	//4. 給定下標
	var numArr04 = [...]int{1: 12, 0: 11, 2: 13}
	fmt.Println("numArr04=", numArr04)
}
