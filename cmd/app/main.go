package main

import (
	"log"

	"github.com/iamthe1whoknocks/pipeline_test_task/config"
	"github.com/iamthe1whoknocks/pipeline_test_task/internal"
)

type App struct {
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("app - config.Load() :%s", err)
	}

	log.Printf("app - loaded config : \n%+v\n", cfg)

	app := internal.NewApp(cfg)

	app.Run()

}
