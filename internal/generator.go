package internal

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/iamthe1whoknocks/pipeline_test_task/config"
)

// data generator model
type Generator struct {
	SendCh      chan []int //chan to send slice to
	SendTime    time.Duration
	SliceLength int
}

// new generator creation
func NewGenerator(cfg *config.Config, ch chan []int) *Generator {
	return &Generator{
		SendCh:      ch,
		SendTime:    cfg.InputTime,
		SliceLength: cfg.InputSliceLength,
	}
}

// generate new slices and send to chan every SendTime ms
func (g *Generator) Run(ctx context.Context) {
	ticker := time.NewTicker(g.SendTime)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Printf("generator - generate - ctx.Done()")
			return
		case <-ticker.C:
			res, err := generateIntSlice(g.SliceLength)
			if err != nil {
				log.Printf("generator - generate - generateIntSlice :%s", err)
				continue
			}
			log.Printf("generator : %s, slice : %v", time.Now().String(), res)
			g.SendCh <- res
		}
	}
}

// generate random int slice
func generateIntSlice(length int) ([]int, error) {
	res := make([]int, 0, length)

	for i := 0; i < length; i++ {
		val, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return nil, fmt.Errorf("generateIntSlice : %w", err)
		}
		res = append(res, int(val.Int64()))
	}
	return res, nil
}
