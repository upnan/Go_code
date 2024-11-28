package main

import (
	"fmt"
)

type Usb interface {
	start()
	stop()
}

type Phone struct {
}

func (p Phone) start() { //实现接口
	fmt.Println("手机工作")
}
func (p Phone) stop() {
	fmt.Println("手机关闭")
}

type Camera struct {
}
//方法属于usb的
func (c Camera) start() {
	fmt.Println("相机工作")
}
func (c Camera) stop() {
	fmt.Println("相机关闭")
}

type Computer struct {
}

func (com Computer) working(usb Usb) {
	usb.start()
	usb.stop()
}

func main() {
	a := Phone{}
	b := Camera{}
	c := Computer{}

	c.working(a)		//传入接口类型，自动识别实现接口的类型，实现多态，类似于c++纯虚函数
	c.working(b)
}
