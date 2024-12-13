package main

import (
	"fmt"
	"log"

	"github.com/anatolygg/tg_bot/internal/config"
	"github.com/anatolygg/tg_bot/internal/logger"
)

func main() {
	cfg, err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		log.Fatal(err.Error())
	}

	log, err := logger.InitLogger(cfg.LogPath, true)
	if err != nil {
		fmt.Println(err.Error())
	}
	_ = log

	token := config.LoadEnv("TOKEN", "")

	_ = token

	// tgClient := telegram.New(token)

	// fetcher := fetcher.New()

	// processor := processor.New()
	// consumer.start(fetcher, processor)

}
