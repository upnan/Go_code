package main

import (
	"fmt"
	"sync"
)

// 一个具有竞争的程序
// var x = 0

// func increment(wg *sync.WaitGroup, m *sync.Mutex) { //传入的是地址值，不能直接值传递，否则锁都是不同的，无效
// 	//解决方法，增加互斥锁变量
// 	m.Lock()
// 	x = x + 1
// 	//m.Unlock()
// 	wg.Done()
// 	//m.Unlock()锁哪里最后结果都一样
// }

// func main() {
// 	var wg sync.WaitGroup
// 	var m sync.Mutex
// 	for i := 0; i < 1000; i++ {
// 		wg.Add(1)
// 		go increment(&wg, &m)
// 	}
// 	wg.Wait()
// 	fmt.Println(x) //最后得出的结果都是随机的，因为在相互竞争
// }

// 使用管道解决竞争问题
var x = 0

func increment(wg *sync.WaitGroup, c chan bool) { //必须引用导入，因为等待结束不能进行复制
	c <- true
	x = x + 1
	<-c
	wg.Done()
}

func main() {
	c := make(chan bool, 1) //只能存一个，第一个进程进去后，其他进程被阻塞，只能等待加一完成
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg, c)
	}
	wg.Wait()
	fmt.Println(x)
}

/*通常，在 goroutine 需要相互通信时使用 通道，
在仅一个 goroutine 应该访问代码的临界部分时使用 互斥锁。
*/
