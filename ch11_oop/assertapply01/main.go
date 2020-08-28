package main

import "fmt"

type usb interface {
	start()
	stop()
}

type phone struct {
	name string
}

func (p phone) start() {
	fmt.Println("手機開始工作")
}

func (p phone) stop() {
	fmt.Println("手機停止工作")
}

func (p phone) call() {
	fmt.Println("手機正在打電話")
}

type camera struct {
	name string
}

func (c camera) start() {
	fmt.Println("相機開始工作")
}

func (c camera) stop() {
	fmt.Println("相機停止工作")
}

type computer struct {
}

func (c computer) Working(usb usb) {
	usb.start()
	if p, ok := usb.(phone); ok {
		p.call()
	}
	usb.stop()
}

func main() {
	var usbArr []usb
	usbArr = append(usbArr, phone{"apple"})
	usbArr = append(usbArr, phone{"asus"})
	usbArr = append(usbArr, camera{"sony"})

	var com computer
	for _, value := range usbArr {
		com.Working(value)
		fmt.Println()
	}
}
