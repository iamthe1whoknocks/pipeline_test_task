package internal

import (
	"context"
	"io"
	"log"
	"sync"
	"time"

	"github.com/iamthe1whoknocks/pipeline_test_task/config"
)

// publicator model
type Publicator struct {
	*Accumulator
	PublTime time.Duration
	Output   io.ReadWriter
}

// new publicator creation
func NewPublicator(cfg *config.Config, acc *Accumulator) *Publicator {
	return &Publicator{
		PublTime:    cfg.PublicateTime,
		Accumulator: acc,
	}
}

// sum values from handler and store it in Value
func (p *Publicator) Run(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	ticker := time.NewTicker(p.PublTime)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			wg.Done()
			log.Printf("publicator - Run - ctx.Done()")
			return
		case <-ticker.C:
			p.Publicate()
		}
	}
}

// send value to source
func (p *Publicator) Publicate() {
	result := p.Accumulator.GetValue()
	log.Printf("publicator - time : %v, result value = %d", time.Now(), result)
}
