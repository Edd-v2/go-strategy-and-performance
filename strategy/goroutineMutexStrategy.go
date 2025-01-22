package strategy

import (
	"fmt"
	"sync"
	"time"
)

type GoroutineMutexStrategy struct{}

func (g *GoroutineMutexStrategy) Execute(workload int) error {
	var mutex sync.Mutex
	for i := 0; i < workload; i++ {
		go func(i int) {
			mutex.Lock()
			defer mutex.Unlock()
			time.Sleep(10 * time.Millisecond)
			fmt.Printf("Task %d completed\n", i)
		}(i)
	}
	time.Sleep(100 * time.Millisecond)
	return nil
}
