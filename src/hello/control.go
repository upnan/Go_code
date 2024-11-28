// //控制语句

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	//if语句
// 	// a := 10
// 	// if a := 1; a > 0 {		//这里的a是局部变量，不影响外面的a
// 	// 	fmt.Println(a)
// 	// 	a++					//只能作为方法使用，不能用作表达式右侧
// 	// 	fmt.Println(a)
// 	// }
// 	// fmt.Println(a)

// 	//for语句
// 	//第一种
// 	// a := 1
// 	// for {
// 	// 	a++
// 	// 	if a > 3 {
// 	// 		break
// 	// 	}
// 	// 	fmt.Println(a)
// 	// }
// 	// fmt.Println("Over!")

// 	//第二种
// 	// a := 1
// 	// for a < 3 {
// 	// 	a++
// 	// 	fmt.Println(a)
// 	// }
// 	// fmt.Println("Over!")

// 	//第三种
// 	// a := 1
// 	// for i := 1; i < 3; i++ {
// 	// 	a++
// 	// 	fmt.Println(a)
// 	// }
// 	// fmt.Println("Over")

// 	//switch语句

// 	// switch a:=1;{			//a是局部变量
// 	// case a>=0:
// 	// 	fmt.Println("a>=0")
// 	// 	fallthrough			//不跳出switch,继续执行
// 	// case a>=1:
// 	// 	fmt.Println("a>=1")
// 	// default:
// 	// 	fmt.Println("None")	
// 	// }

// 	//跳转语句
// 	// LABLE1:
// 		for{
// 			for i := 0; i < 10; i++ {
// 				if(i>3){
// 					goto LABLE1		//跳出标签整体,类似有goto跳到标签处，continue是跳过该次循环，不执行下面的语句了
// 				}
// 			}
// 		}
// 	LABLE1:
// }
