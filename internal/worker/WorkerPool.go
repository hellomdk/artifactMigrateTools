package worker

import (
	"sync"
)

type WorkerPool struct {
	workerNum   int
	taskQueue   chan Task
	stopChan    chan struct{}
	workerGroup sync.WaitGroup
}

func NewWorkerPool(workerNum int) *WorkerPool {
	return &WorkerPool{
		workerNum: workerNum,
		taskQueue: make(chan Task),
		stopChan:  make(chan struct{}),
	}
}

func (p *WorkerPool) Run() {
	for i := 0; i < p.workerNum; i++ {
		go p.workerFunc()
		p.workerGroup.Add(1)
	}
}

func (p *WorkerPool) workerFunc() {
	defer p.workerGroup.Done()
	for {
		select {
		case task := <-p.taskQueue:
			task.Handler(task.Args)
		case <-p.stopChan:
			return
		}
	}
}

func (p *WorkerPool) Stop() {
	close(p.stopChan)
	p.workerGroup.Wait()
}

func (p *WorkerPool) AddTask(task Task) {
	p.taskQueue <- task
}
