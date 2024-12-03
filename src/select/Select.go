package main

import (
	"fmt"
	"time"
)

// select实例
func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "server from 1"
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "server from 2"
}

// func main() {
// 	ch1 := make(chan string)
// 	ch2 := make(chan string)
// 	go server1(ch1)
// 	go server2(ch2)
// 	select {		//此处会阻塞，知道一个case完成再结束
// 	case s1 := <-ch1:
// 		fmt.Println(s1)
// 	case s2 := <-ch2:
// 		fmt.Println(s2)
// 	}
// }

// default case 默认case，在其他case未就绪时，使用该case
func process(ch chan string) {
	time.Sleep(10050 * time.Millisecond)
	ch <- "process succeed"
}

// func main() {
// 	ch := make(chan string)
// 	go process(ch)
// 	for {
// 		time.Sleep(1000 * time.Millisecond)
// 		select {
// 		case s1 := <-ch:
// 			fmt.Println(s1)
// 			return
// 		default: //出现default就不会出现死锁问题
// 			fmt.Println("no")
// 		}
// 	}
// }

// 随机打印，若select上case全部就绪，则随机进行一个
func server3(ch chan string) {
	ch <- "server3"
}
func server4(ch chan string) {
	ch <- "server4"
}

func main() {
	ch3 := make(chan string)
	ch4 := make(chan string)

	go server3(ch3)
	go server4(ch4)

	select {
	case s3 := <-ch3:
		fmt.Println(s3)
	case s4 := <-ch4:
		fmt.Println(s4) //输出结果均是随机的
	}
}
