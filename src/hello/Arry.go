// package main

// import (
// 	"fmt"
// )

// func main() {
	//var a [2]int
	// a := [2]int{1, 2}
	// a := [20]int{19: 1}//索引为19为1
	// a := [...]int{1, 2, 3, 4, 5}     //不定义具体个数，计算机自己决定

	// a := [...]int{99: 1}
	// var p *[100]int = &a		//指向数组的指针

	// x, y := 1, 2
	// a := [...]*int{&x, &y}		//指针数组
	// fmt.Println(a)

	// p := new([10]int) //指向数组的指针
	// p[1] = 2          //可直接用索引对该数组指针操作
	// fmt.Println(p)

	//二维数组
	// a := [2][3]int{{1, 2, 3}, {2, 3, 4}}
	// fmt.Println(a)

	//冒泡排序
// 	a := [...]int{3, 2, 6, 7, 9}
// 	fmt.Println(a)
// 	num := len(a)
// 	for i := 0; i < num; i++ {
// 		for j := i + 1; j < num; j++ {
// 			if a[i] > a[j] {
// 				temp := a[i]
// 				a[i] = a[j]
// 				a[j] = temp
// 			}
// 		}
// 	}
// 	fmt.Println(a)
// }
