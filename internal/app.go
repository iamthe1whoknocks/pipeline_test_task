package internal

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/iamthe1whoknocks/pipeline_test_task/config"
)

type App struct {
	Config *config.Config
}

func NewApp(cfg *config.Config) *App {
	return &App{
		Config: cfg,
	}
}

func (a *App) Run() {
	ctx, cancel := context.WithCancel(context.Background())

	wg := sync.WaitGroup{}
	genCh := make(chan []int)
	accumCh := make(chan []int)

	gen := NewGenerator(a.Config, genCh)
	go gen.Run(ctx, &wg)

	handler := NewHandler(a.Config, genCh, accumCh)
	go handler.Run(ctx, &wg)

	acc := NewAccumulator(accumCh)
	go acc.Run(ctx, &wg)

	publicator := NewPublicator(a.Config, acc)
	go publicator.Run(ctx, &wg)

	// graceful shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		cancel()
		wg.Wait()
		log.Printf("main - Run - signal: %s", s.String())
	}
}
