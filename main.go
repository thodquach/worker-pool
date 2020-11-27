package main

import (
	"fmt"
	"time"

	"worker-pool/pool"
	"worker-pool/task"
)

func main() {
	var tasks []*task.Task
	p := pool.NewPool(tasks, 3)
	p.Start()

	for i := 0; i < 10; i++ {
		element := i
		currentTask := task.NewTask(func() error {
			fmt.Printf("The nth: %d element task run in workerpool. \n", element)
			time.Sleep(3 * time.Second)
			return nil
		})
		p.RunSingleTask(currentTask)
	}

	p.Stop()
}
