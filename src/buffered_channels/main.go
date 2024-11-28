package main

// import (
// 	"fmt"
// 	//"time"
// )

// 缓冲通道的创建
// func main() {
// 	a := make(chan int, 2)
// 	a <- 2 //因为有缓冲区，不会被阻塞
// 	a <- 3
// 	fmt.Println(<-a)
// 	fmt.Println(<-a)
// }

// func write(ch chan int) {
// 	for i := 0; i < 5; i++ {
// 		ch <- i
// 		fmt.Println("写入", i)
// 	}
// 	close(ch)
// }

// func main() {
// 	ch := make(chan int, 2)
// 	go write(ch)
// 	time.Sleep(2 * time.Second)//此时休眠，缓冲通道写满
// 	for v := range ch {
// 		fmt.Println("read", v)	//此处读一个写一个
// 		time.Sleep(2 * time.Second)//只要主进程休眠，便会跑至子进程运行，直至子进程完成
// 	}
// }

// 关闭缓冲通道
// func main() {
// 	ch := make(chan int, 2)
// 	ch <- 1
// 	ch <- 2
// 	close(ch)

// 	n, open := <-ch	//存在值的话，返回通道中的值，并返回true
// 	fmt.Println(n, open)
// 	n, open = <-ch
// 	fmt.Println(n, open)
// 	n, open = <-ch
// 	fmt.Println(n, open)//不存在返回0，和false
// }

// 缓冲通道的容量（初始订的，不变）和长度(所排列的元素)
// func main() {
// 	ch := make(chan string, 3)
// 	ch <- "naveen"
// 	ch <- "paul"
// 	fmt.Println("capacity is", cap(ch))
// 	fmt.Println("length is", len(ch))
// 	fmt.Println("read value", <-ch)
// 	fmt.Println("new length is", len(ch))
// }

//waitGroup,实现完成多个子Goroutine，后完成主Goroutine
// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func process(i int, wg *sync.WaitGroup) {
// 	fmt.Println("Goroutine", i, "start")
// 	time.Sleep(2 * time.Second)//该等待期间，所有子goroutine完成了自己的前半部分
// 	fmt.Println("Goroutine", i, "ended")
// 	wg.Done() //将计数器减一
// }

// func main() {
// 	var wp sync.WaitGroup
// 	for i := 0; i < 3; i++ {
// 		go process(i, &wp)
// 		wp.Add(1)//使计数器加1
// 	}//该循环创建了三个子Goroutine
// 	wp.Wait()//等待所有子Goroutine完成
// 	fmt.Println("主Goroutine完成")
// }

//worker pool实现

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 创建表示作业和结果的结构体
type Job struct {
	id       int
	randomno int
}
type Result struct {
	job         Job //对应作业
	sumofdigits int //计算结果
}

// 创建用于接收作业，和写入输出的缓冲通道
var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

// 计算数字各位之和函数
func digits(number int) int {
	sum := 0
	for number != 0 {

		sum += number % 10
		number /= 10
	}
	time.Sleep(2 * time.Second) //用于模拟计算所要时间
	return sum
}

// 编写一个函数来创建一个 worker Goroutine。（寻找是否有作业，存在就计算该作业，并存入结果缓冲通道）
func worker(wp *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wp.Done()
}

// createWorkerPool 函数将创建一个 worker Goroutines 池。
func createWorkerPool(noOfworkers int) { //传入的是工作线程数量
	var wp sync.WaitGroup
	for i := 0; i < noOfworkers; i++ {
		go worker(&wp) //传入引用，而不是新创建的副本
		wp.Add(1)
	}
	wp.Wait()
	close(results) //在所有 Goroutines 完成执行后，它会关闭 results 通道，
	//因为所有 Goroutines 都已完成执行，没有其他人会进一步写入 results 通道。
}

// 创建多个所要处理的作业
func allocate(noOfjobs int) {
	for i := 0; i < noOfjobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

// 打印结果
func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true //输出完所有结果写入
}

func main() {
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 10
	createWorkerPool(noOfWorkers) //开始工作，创造多个线程
	<-done                        //设立此处，达到没输出完就一直休眠的效果
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}

//完成该工作吃池很累
