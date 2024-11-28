// package main

// import (
// 	"fmt"
// )

// type A struct {
// 	name string
// }

// type B struct {
// 	Name string
// }

// type Tz int

// func main() {
// 	// var a A
// 	// a.name = "jack"
// 	// fmt.Println(a.name)
// 	// a.Print()
// 	// fmt.Println(a.name)

// 	// var b B
// 	// b.Print()
// 	// fmt.Println(b.Name)
// 	var a Tz
// 	a.add_100()
// 	fmt.Println(a)
// }

// func (a *Tz) add_100() {
// 	*a += 100
// }

// func (a *A) Print() { //方法，将类型与函数相绑定，可以是任何类型，增加灵活
// 	a.name = "jack" //相当于class中方法的定义，可以访问结构体中私有字段
// 	fmt.Println("A")
// }

// func (b B) Print() {
// 	b.Name = "mark"
// 	fmt.Println("B")
// }
