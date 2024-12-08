package main

import (
	"fmt"
	"sync"
)

type rect struct {
	length int
	width  int
}

func (r rect) area(wg *sync.WaitGroup) { //结构体rect的area方法
	defer wg.Done() //该处使用defer使函数最后结束都会对wg减1操作，不用每次判断结束后都进行减一
	if r.length < 0 {
		fmt.Println("eof")
		return
	}
	if r.width < 0 {
		fmt.Println("eof")
		return
	}
	area := r.length * r.width
	fmt.Println(area)
}

func main() {
	var wg sync.WaitGroup
	r1 := rect{-67, 89}
	r2 := rect{5, -67}
	r3 := rect{8, 9}
	rects := []rect{r1, r2, r3}
	for _, v := range rects {
		wg.Add(1)
		go v.area(&wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}
