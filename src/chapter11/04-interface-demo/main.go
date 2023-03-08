package main

import "fmt"

// Usb 聲明一個介面
type Usb interface {
	//聲明了兩個沒有實現的方法
	Start()
	Stop()
}

type Phone struct {
}

// 讓Phone實現Usb的方法
func (p Phone) Start() {
	fmt.Println("手機開始工作...")
}
func (p Phone) Stop() {
	fmt.Println("手機停止工作...")
}

type Camera struct {
}

// 讓Camera實現Usb的方法
func (c Camera) Start() {
	fmt.Println("相機開始工作...")
}
func (c Camera) Stop() {
	fmt.Println("相機停止工作...")
}

type Computer struct {
}

// Working 編寫一個方法，方法接收一個Usb介面類型
// 只要實現Usb介面，所謂實現Usb介面，就是指實現了Usb介面聲明的所有方法
func (c Computer) Working(usb Usb) {
	//透過Usb介面變數來調用Start和Stop方法
	usb.Start()
	usb.Stop()
}

func main() {
	computer := Computer{}

	phone := Phone{}
	camera := Camera{}

	computer.Working(phone)
	computer.Working(camera)
}
