// package main

// import (
// 	"fmt"
// )

// func main() {
// 	f := closeure(10)
// 	fmt.Println(f(1))
// 	fmt.Println(f(2))
// 	A()
// 	B()
// 	C()
// }

// // 闭包
// func closeure(x int) func(int) int { //用的都是同一个x，对应了地址值
// 	fmt.Printf("%p\n", &x)
// 	return func(y int) int {
// 		fmt.Printf("%p\n", &x)
// 		return x + y
// 	}
// }

// //defer,recover,...等异常处理

// func A() {
// 	fmt.Println("Func A")
// }

// func B() {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println("Recover in B")
// 		}
// 	}()
// 	panic("Panic In B")
// }

// func C() {
// 	fmt.Println("Func In C")
// }
