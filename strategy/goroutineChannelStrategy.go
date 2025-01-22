package strategy

import (
	"fmt"
	"time"
)

type GoroutineChannelStrategy struct{}

func (g *GoroutineChannelStrategy) Execute(workload int) error {
	ch := make(chan int)

	for i := 0; i < workload; i++ {
		go func(i int) {
			time.Sleep(10 * time.Millisecond)
			ch <- i
			fmt.Printf("Task %d completd\n", i)
		}(i)
	}

	for i := 0; i < workload; i++ {
		<-ch
	}
	return nil
}
