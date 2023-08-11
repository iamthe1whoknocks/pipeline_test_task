package internal

import (
	"context"
	"log"
	"sort"
	"time"

	"github.com/iamthe1whoknocks/pipeline_test_task/config"
)

// handler model
type Handler struct {
	ReceiveCh         <-chan []int
	SendCh            chan<- []int
	OutputSliceLength int
	InputSliceLength  int
	WorkerNum         int
}

// new handler creation
func NewHandler(cfg *config.Config, receiveCh, sendCh chan []int) *Handler {
	return &Handler{
		SendCh:            sendCh,
		ReceiveCh:         receiveCh,
		OutputSliceLength: cfg.OutputSliceLength,
		InputSliceLength:  cfg.InputSliceLength,
		WorkerNum:         cfg.WorkerNum,
	}
}

// Handling incoming data from generator and send to SendCh
func (h *Handler) Run(ctx context.Context) {
	for w := 1; w <= h.WorkerNum; w++ {
		go h.worker(w)
	}

	for {
		select {
		case <-ctx.Done():
			log.Printf("handler - handle - ctx.Done()")
			return
		default:
			// case data := <-h.ReceiveCh:
			// 	handleSlice(data, h.InputSliceLength, h.OutputSliceLength)
		}
	}
}

func (h *Handler) worker(id int) {
	for j := range h.ReceiveCh {
		log.Printf("handler - worker№%d - input : %v\n", id, j)
		time.Sleep(time.Second * 2) // for clarity
		sort.Ints(j)
		out := j[h.InputSliceLength-h.OutputSliceLength:]
		log.Printf("handler - worker№%d - output : %v\n", id, out)
		h.SendCh <- out
	}
}
