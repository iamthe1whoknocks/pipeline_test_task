package main

import (
	"context"
	"log"

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

	genCh := make(chan []int)
	accumCh := make(chan []int)

	gen := internal.NewGenerator(cfg, genCh)
	go gen.Run(context.TODO())

	handler := internal.NewHandler(cfg, genCh, accumCh)
	go handler.Run(context.TODO())

	acc := internal.NewAccumulator(cfg, accumCh)
	go acc.Run(context.TODO())

}
