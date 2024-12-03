package main

import (
	"fmt"
	"time"
)

// selectÊµÀý
func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "server from 1"
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "server from 2"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go server1(ch1)
	go server2(ch2)
	select {
	case s1 := <-ch1:
		fmt.Println(s1)
	case s2 := <-ch2:
		fmt.Println(s2)
	}
}
