package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("hello")
}

func dig() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func abc() {
	for i := 'a'; i < 'z'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

/*
1.当启动新的 Goroutine 时，goroutine 调用会立即返回。
与函数不同，控制权不会等待 Goroutine 执行完毕。
在 Goroutine 调用之后，控制权会立即返回到下一行代码，并且会忽略 Goroutine 的任何返回值。
2.主 Goroutine 必须处于运行状态，其他 Goroutine 才能运行。如果主 Goroutine 终止，则程序将终止，并且不会运行其他 Goroutine。
*/
func main() {
	// go hello()
	// time.Sleep(1 * time.Second) //是主线程休眠，给子线程足够时间输出
	// fmt.Println("main function")

	go dig()
	go abc() //证明同时进行，在主线程休眠的过程中，两个线程交替进行

	time.Sleep(3000 * time.Millisecond)

	fmt.Printf("main function")
}
