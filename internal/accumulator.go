package internal

import (
	"context"
	"log"
	"sync"

	"github.com/iamthe1whoknocks/pipeline_test_task/config"
)

// accumulator model
type Accumulator struct {
	ReceiveCh <-chan []int
	Mu        *sync.RWMutex
	Value     int
}

// new accumulator creation
func NewAccumulator(cfg *config.Config, receiveCh chan []int) *Accumulator {
	return &Accumulator{
		ReceiveCh: receiveCh,
		Mu:        new(sync.RWMutex),
		Value:     0,
	}
}

// sum values from handler and store it in Value
func (a *Accumulator) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Printf("handler - handle - ctx.Done()")
			return
		case data := <-a.ReceiveCh:
			var sum int
			for _, v := range data {
				sum += v
			}
			a.Mu.Lock()
			a.Value += sum
			a.Mu.Unlock()
		}
	}
}
