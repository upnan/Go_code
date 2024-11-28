package main

import (
	"fmt"
)

// // 创建管道
// // func main() {
// // 	var a chan int

// // 	if a==nil  {
// // 		fmt.Println("make it")
// // 		a = make(chan int)
// // 		fmt.Printf("%T", a)
// // 	}
// // }

// //默认情况下，管道的发送和接受都是阻塞的。
// //当数据发送到通道时，控制在发送语句中被阻塞，直到其他 Goroutine 从该通道读取。
// //利用此特性实现进程之间的并发

// func hello(a chan bool) {
// 	fmt.Println("hello function")
// 	a <- true //写入管道，（通道箭头只有一个方向）
// }

// // func main() {
// // 	a := make(chan bool)
// // 	go hello(a) //进行此时，立即跳转到下一行
// // 	<-a         //管道接受数据，阻塞此处，此时进行hello（优先主进程，但主进程休眠，启动子进程）
// // 	fmt.Println("main function")
// // }

// //求一个数各位平方和和立方和之和

// func calcSquares(a int, squares chan int) {
// 	sum := 0
// 	for a != 0 {
// 		digit := a % 10
// 		a /= 10
// 		sum = sum + digit*digit
// 	}
// 	fmt.Println("squares")
// 	squares <- sum //写入管道
// }

// func clacCubes(a int, cubes chan int) {
// 	sum := 0
// 	for a != 0 {
// 		digit := a % 10
// 		a /= 10
// 		sum += digit * digit * digit
// 	}
// 	fmt.Println("cubes")
// 	cubes <- sum
// }

// func main() {
// 	squares := make(chan int)
// 	cubes := make(chan int)
// 	go calcSquares(589, squares)
// 	go clacCubes(589, cubes)	//先运行下面的进程，再是上面的
// 	square, cube := <-squares, <-cubes
// 	fmt.Println(square + cube)
// }
/*clacCubes 先运行的现象是并发程序中常见的，由 Go 调度器决定的执行顺序。
它并不一定每次都会是立方的计算先执行，其他的运行可能是平方计算先执行。
在并发程序中，执行顺序的变化是正常现象，且通常不影响程序的功能，只要最终的结果正确返回。*/

//单向通道
// func a(c chan<-int){
// 	c<-10
// }

// func main(){
// 	c:=make(chan<-int)
// 	go a(c)
// 	fmt.Println(<-c)//单向通道不能接收，会报错
// }

// 单向通道的作用（临时转换为单向）
// func a(c chan<- int) {
// 	c <- 10
// }

// // func main() {
// // 	c := make(chan int)//将双向的管道当作单向管道传入，在该线程中作为单向管道，退出后仍为双向管道
// // 	go a(c)
// // 	fmt.Println(<-c)
// // }

// func producer(chnl chan int) {
// 	for i := 1; i <= 9; i++ {
// 		chnl <- i //生产者阻塞，等待消费者接收
// 	}
// 	close(chnl) //关闭通道
// }
// func main() {
// 	chnl := make(chan int)
// 	go producer(chnl)
// 	//for {
// 	// v, ok := <-chnl//ok识别到管道关闭，会赋值为false。若不关闭管道，会发生死锁，该处已经停止接收
// 	// //而如果没有关闭通道，消费者就永远无法知道生产者是否已经完成了所有的数据发送，从而可能导致死锁。
// 	// if ok == false {
// 	// 	break
// 	// }

// 	//第2种写法
// 	for v := range chnl {
// 		fmt.Printf("receved %d\n", v)
// 	}
// 	//}
// }

//重写平方和立方程序，利用管道存储数字的各位

func digits(number int, dig chan int) {
	for number != 0 {
		digit := number % 10
		dig <- digit
		number /= 10
	}
	close(dig) //关闭通道，防止死锁
}

func caluSqurse(number int, squres chan int) {
	dig := make(chan int)
	go digits(number, dig)
	sum := 0
	for v := range dig {
		sum += v * v
	}
	squres <- sum
}

func caluCubes(number int, cubes chan int) {
	dig := make(chan int)
	go digits(number, dig)
	sum := 0
	for v := range dig {
		sum += v * v * v
	}
	cubes <- sum
}

func main() {
	number := 589
	squres := make(chan int)
	cubes := make(chan int)
	go caluSqurse(number, squres)
	go caluCubes(number, cubes)
	sq, cu := <-squres, <-cubes
	fmt.Println(sq + cu)
}
