package helper

import (
	"fmt"
	"go-strategy-and-performance/strategy"
)

// Mappa delle strategie
var strategies = map[string]strategy.Strategy{
	"sequential":        &strategy.SequentialStrategy{},
	"goroutine_channel": &strategy.GoroutineChannelStrategy{},
	"goroutine_mutex":   &strategy.GoroutineMutexStrategy{},
}

func GetStrategy(strategy_name string) (strategy.Strategy, error) {
	strategy, exists := strategies[strategy_name]
	if !exists {
		return nil, fmt.Errorf("invalid strategy: %s", strategy_name)
	}
	return strategy, nil
}
