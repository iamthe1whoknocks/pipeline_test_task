package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/iamthe1whoknocks/pipeline_test_task/config"
	"github.com/iamthe1whoknocks/pipeline_test_task/internal"
)

type App struct {
	Generator *internal.Generator
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("app - config.Load() :%s", err)
	}

	log.Printf("app - loaded config : \n%+v\n", cfg)

	ctx, cancel := context.WithCancel(context.Background())

	wg := sync.WaitGroup{}
	genCh := make(chan []int)
	accumCh := make(chan []int)

	gen := internal.NewGenerator(cfg, genCh)
	go gen.Run(ctx, &wg)

	handler := internal.NewHandler(cfg, genCh, accumCh)
	go handler.Run(ctx, &wg)

	acc := internal.NewAccumulator(accumCh)
	go acc.Run(ctx, &wg)

	publicator := internal.NewPublicator(cfg, acc)
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
