// package main

// import (
// 	"fmt"
// )

// const (
// 	a = iota
// 	b = iota //1.遇到const关键字会被重置为0；2.const组中，每增加一个变量，iota加1;3.只能在const组中用
// 	c = iota
// 	d = iota
// )

// // 实现计算机储存单位的枚举
// const (
// 	B float64 = 1 << (iota * 10) 
// 	KB
// 	MB
// 	GB
// )

// func main() {
// 	fmt.Println(B)
// 	fmt.Println(KB)
// 	fmt.Println(MB)
// 	fmt.Println(GB)
// }
