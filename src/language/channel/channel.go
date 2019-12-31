package main

import (
	fmt2 "fmt"
	"time"
)

func worker(id int, c chan int) {
	for n := range c {
		//n, ok := <- c	//判断是否关闭 方法一
		//if !ok {
		// break;
		//}

		fmt2.Printf("Worker %d get %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int) //无缓冲通道
	go worker(id, c)
	return c
}

func chanDemo() {
	var channels [10]chan<- int
	for k, _ := range channels {
		channels[k] = createWorker(k)
	}

	for k, _ := range channels {
		channels[k] <- 'a' + k
	}
	time.Sleep(time.Microsecond) //加它的原因式不知道另外一个协程是否打印完毕，只能等待一段时间，这段时间内就算没打印完毕也不管
}

func bufferedChannel() {
	c := make(chan int, 3) //有缓冲通道
	go worker(0, c)
	c <- 'a'
	c <- 'A'
	c <- 'b'
	c <- 'B'
	time.Sleep(time.Microsecond)
}

func channelClose() {
	//发送方close 通知 接收方
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'A'
	c <- 'b'
	c <- 'B'
	close(c)
}

func main() { //goroutine
	//chanDemo()	//无缓冲通道

	//bufferedChannel()	//有缓冲通道

	//channelClose()	//关闭通道
}
