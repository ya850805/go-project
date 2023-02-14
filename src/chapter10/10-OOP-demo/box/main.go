package main

import "fmt"

type Box struct {
	length float64
	width  float64
	height float64
}

func (b *Box) getVolume() float64 {
	return b.length * b.width * b.height
}

func main() {
	//1. 編寫創建一個`Box`結構體，在其中聲明三個字段表示立方體的長、寬、高，長寬高要從終端獲取。
	//2. 聲明一個方法獲取立方體的體積。
	//3. 創建一個`Box`結構體變數，打印給定尺寸的立方體的體積。
	var box Box

	fmt.Print("請輸入盒子的長：")
	fmt.Scanln(&box.length)
	fmt.Print("請輸入盒子的寬：")
	fmt.Scanln(&box.width)
	fmt.Print("請輸入盒子的高：")
	fmt.Scanln(&box.height)

	volume := box.getVolume()
	fmt.Printf("體積為：%.2f", volume)
}
