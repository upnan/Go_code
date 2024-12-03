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

func main() {
	ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case s1 := <-ch:
			fmt.Println(s1)
			return
		default:
			fmt.Println("no")
		}
	}
}
