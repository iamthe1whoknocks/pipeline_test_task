package config

import (
	"os"
	"strconv"
	"time"
)

// env
const (
	InputSliceLength        = 10 // int slice length in incoming data
	OutputSliceLength       = 3  // int slice length in incoming data
	EnvInputTimeMs          = "PIPELINE_NEW_INPUT_TIME_MS"
	EnvWorkerNum            = "PIPELINE_WORKER_NUM"
	EnvPublicatorOutputTime = "PIPELINE_PUBLICATOR_OUTPUT_TIME_SEC"
)

// config model
type Config struct {
	InputTime         time.Duration // incoming data packets frequency
	WorkerNum         int           // number of workers use
	PublicateTime     time.Duration // time to publicate results frequency
	InputSliceLength  int
	OutputSliceLength int
}

// load env variables to config
func Load() (*Config, error) {
	cfg := &Config{}

	updateTimeInt, err := strconv.Atoi(os.Getenv(EnvInputTimeMs))
	if err != nil {
		return nil, err
	}
	cfg.InputTime = time.Duration(updateTimeInt) * time.Millisecond

	workerNum, err := strconv.Atoi(os.Getenv(EnvWorkerNum))
	if err != nil {
		return nil, err
	}

	cfg.WorkerNum = workerNum

	publicatorTimeInt, err := strconv.Atoi(os.Getenv(EnvPublicatorOutputTime))
	if err != nil {
		return nil, err
	}
	cfg.PublicateTime = time.Duration(publicatorTimeInt) * time.Second

	cfg.InputSliceLength = InputSliceLength
	cfg.OutputSliceLength = OutputSliceLength

	return cfg, nil
}
