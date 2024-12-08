package main

import (
	"fmt"
	"time"
)

func tatalTime(start time.Time) {
	fmt.Printf("Total time taken: %f seconds", time.Since(start).Seconds()) //输出开始时间与当前时间的差值
}

func test() {
	start := time.Now()
	defer tatalTime(start) //用于计算出test（）函数使用的总时间,在test返回之前执行该函数
	time.Sleep(2 * time.Second)
	fmt.Println("test completed")
}

func displayTest(a int) {
	fmt.Println(a)
}

// func main() {
// 	//test()
// 	// a := 5
// 	// defer displayTest(a) //参数在执行之前就已经确定，后面发生变化也不会更改
// 	// a = 10
// 	// fmt.Println(a)

// 	//多个推迟函数，最后按照后进先出结束
// 	str := "abcdefg"

// 	for _, c := range str {
// 		defer fmt.Printf("%c", c) //实现反向输出字符串
// 	}
// }
