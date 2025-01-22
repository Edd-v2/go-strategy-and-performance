package strategy

import (
	"fmt"
	"time"
)

type SequentialStrategy struct{}

func (s *SequentialStrategy) Execute(workload int) error {
	for i := 0; i < workload; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("Task %d completed\n", i)
	}
	return nil
}
