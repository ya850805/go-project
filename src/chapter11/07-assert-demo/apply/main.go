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
func (p Phone) call() {
	fmt.Println("手機打電話...")
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

type Computer struct {
}

func (computer Computer) Working(usb Usb) {
	usb.Start()

	//如果usb是Phone結構體變數，則再調用call()
	phone, ok := usb.(Phone)
	if ok {
		phone.call()
	}

	usb.Stop()
}

func main() {
	var usbArr [3]Usb

	usbArr[0] = Phone{"phone1"}
	usbArr[1] = Phone{"phone2"}
	usbArr[2] = Camera{"camera1"}

	var computer Computer

	for _, usb := range usbArr {
		computer.Working(usb)
		fmt.Println()
	}
}
