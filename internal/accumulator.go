package internal

import (
	"context"
	"log"
	"sync"
)

// accumulator model
type Accumulator struct {
	ReceiveCh <-chan []int
	Mu        *sync.RWMutex
	Value     int
}

// new accumulator creation
func NewAccumulator(receiveCh chan []int) *Accumulator {
	return &Accumulator{
		ReceiveCh: receiveCh,
		Mu:        new(sync.RWMutex),
		Value:     0,
	}
}

// sum values from handler and store it in Value
func (a *Accumulator) Run(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	for {
		select {
		case <-ctx.Done():
			log.Printf("accumulator - Run - ctx.Done()")
			wg.Done()
			return
		case data, ok := <-a.ReceiveCh:
			if ok {
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

}

// get actual value from accumulator
func (a *Accumulator) GetValue() int {
	var result int
	a.Mu.RLock()
	result = a.Value
	a.Mu.RUnlock()
	return result
}
