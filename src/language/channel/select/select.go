package main

import (
	"fmt"
	"math/rand"
	"time"
)

func createChannel() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Duration(1500) * time.Millisecond)
		fmt.Printf("Worker %d get %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int) //无缓冲通道
	go worker(id, c)
	return c
}

func main() {
	c1 := createChannel()
	c2 := createChannel()
	worker := createWorker(0)

	var values []int // 消费者和生产者速度不一致，需要缓存起来
	tm := time.After(10 * time.Second)

	tick := time.Tick(time.Second)
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}

		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("len(values)=", len(values))
		case <-tm:
			fmt.Println("bye bye")
			return
		}
	}
}
