package main

import (
	"context"
	"log"
	"time"

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

	log.Printf("app - loaded config : %+v\n", cfg)

	genCh := make(chan []int)
	handlerCh := make(chan []int)
	gen := internal.NewGenerator(cfg, genCh)

	go gen.Generate(context.TODO())

	// for res := range ch {
	// 	log.Printf("main - ch : %v", res)
	// }

	handler := internal.NewHandler(cfg, genCh, handlerCh)

	go handler.Handle(context.TODO())

	time.Sleep(15 * time.Second)

}
