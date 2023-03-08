package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	Name string
}

func (p Phone) Start() {
	fmt.Println("手機開始工作...")
}
func (p Phone) Stop() {
	fmt.Println("手機停止工作...")
}

type Camera struct {
	Name string
}

func (c Camera) Start() {
	fmt.Println("相機開始工作...")
}
func (c Camera) Stop() {
	fmt.Println("相機停止工作...")
}

func main() {
	//定義一個Usb介面陣列，可以存放Phone和Camera結構體變數
	//這裡就體現多行陣列
	var usbArr [3]Usb

	usbArr[0] = Phone{"phone1"}
	usbArr[1] = Phone{"phone2"}
	usbArr[2] = Camera{"camera1"}

	fmt.Println("usbArr=", usbArr)
}
