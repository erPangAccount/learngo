package main

import (
	fmt2 "fmt"
	"sync"
)

const _WORKER_NUM = 20

func doWork(id int, w worker) {
	for n := range w.in {
		fmt2.Printf("Worker %d get %d\n", id, n)
		w.done()
	}
}

type worker struct {
	in   chan int
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w)
	return w
}

func chanDemo() {
	var workers [_WORKER_NUM]worker
	var wg sync.WaitGroup

	for k, _ := range workers {
		workers[k] = createWorker(k, &wg)
	}

	for k, worker := range workers {
		wg.Add(1)
		worker.in <- 'a' + k
	}

	for k, worker := range workers {
		wg.Add(1)
		worker.in <- 'A' + k
	}

	wg.Wait()
}

func main() { //goroutine
	chanDemo() //无缓冲通道
}
