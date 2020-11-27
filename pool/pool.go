package pool

import (
	"sync"

	"worker-pool/task"
)

/*** Pool ***/

// Pool is a worker group that runs a number of tasks at a
// configured concurrency.
type Pool struct {
	Tasks []*task.Task

	concurrency int
	tasksChan   chan *task.Task
	wg          sync.WaitGroup
}

// NewPool initializes a new pool with the given tasks and
// at the given concurrency.
func NewPool(tasks []*task.Task, concurrency int) *Pool {
	return &Pool{
		Tasks:       tasks,
		concurrency: concurrency,
		tasksChan:   make(chan *task.Task),
	}
}

// Only start worker with number process equal to `concurrency`
func (p *Pool) Start() {
	for i := 0; i < p.concurrency; i++ {
		go p.Work()
	}
}

// Add single task into Task channel and do not wait to run concurrent
func (p *Pool) RunSingleTask(singleTask *task.Task) {
	p.wg.Add(1)
	p.tasksChan <- singleTask
	// p.wg.Wait()
}

// Close Task channel and wait to all tasks to be done
func (p *Pool) Stop() {
	// All workers return
	close(p.tasksChan)

	p.wg.Wait()
}

// Run runs all work within the pool and blocks until it's finished.
func (p *Pool) Run() {
	for i := 0; i < p.concurrency; i++ {
		go p.Work()
	}

	p.wg.Add(len(p.Tasks))
	for _, task := range p.Tasks {
		p.tasksChan <- task
	}

	// All workers return
	close(p.tasksChan)

	p.wg.Wait()
}

// The work loop for any single goroutine.
func (p *Pool) Work() {
	for task := range p.tasksChan {
		task.Run(&p.wg)
	}
}
