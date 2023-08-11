package internal

import (
	"context"
	"log"
	"sort"

	"github.com/iamthe1whoknocks/pipeline_test_task/config"
)

// handler model
type Handler struct {
	ReceiveCh         <-chan []int
	SendCh            chan<- []int
	OutputSliceLength int
	InputSliceLength  int
}

// new generator creation
func NewHandler(cfg *config.Config, receiveCh, sendCh chan []int) *Handler {
	return &Handler{
		SendCh:            sendCh,
		ReceiveCh:         receiveCh,
		OutputSliceLength: cfg.OutputSliceLength,
		InputSliceLength:  cfg.InputSliceLength,
	}
}

// Handling incoming data from generator and send to SendCh
func (h *Handler) Handle(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Printf("handler - handle - ctx.Done()")
			return
		case data := <-h.ReceiveCh:
			handleSlice(data, h.InputSliceLength, h.OutputSliceLength)
		}
	}
}

func handleSlice(slice []int, InputSliceLength, outputSliceLength int) []int {
	log.Printf("handler - HandleSlice - input : %v\n", slice)
	sort.Ints(slice)
	log.Printf("handler - HandleSlice - output : %v\n", slice[InputSliceLength-outputSliceLength:])
	return slice[InputSliceLength-outputSliceLength:]
}
