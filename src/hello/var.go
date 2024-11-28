// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// // 全局变量的声明
// // var (
// // 	a    = "aaaa"
// // 	b, c = 1, 2
// // )

// func main() {
// 	//var a int //变量默认值一般为0，bool类型为false，string类型为空字符串
// 	//var a [2]int //容量为2的数组

// 	//变量声明与赋值
// 	//var a int = 10
// 	//类型推断，由系统判断类型
// 	//var a = 10
// 	//更简洁
// 	//a := 1
// 	//a := true

// 	// var a, b, d, e int
// 	// a, b, d, e = 1, 2, 3, 4

// 	// a, _, c, d := 1, 2, 3, 4//下划线代表一个变量，可通过此来忽略此值，常用于函数返回值

// 	// fmt.Println(a, c, d)

// 	//不存在隐式转换，所以要强制；“，”在变量已经创建后可省略；不兼容类型之间不能强制转换，如bool和int
// 	// var b float32 = 100.1
// 	// fmt.Println(b)
// 	// a := int(b)
// 	// fmt.Println(a)

// 	// var a int = 65
// 	// b := string(a)//将数据转化为文本格式，存储的任何东西都为数字，数字65在文本格式表示为A；
// 	// fmt.Println(b)

// 	//如果真的想输出65文本格式,导入strconv包
// 	var a int = 65
// 	b := strconv.Itoa(a)//int表面值转string
// 	a, _ = strconv.Atoi(b)//string转int
// 	fmt.Println(a)
// }
