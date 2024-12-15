package main

import (
	"fmt"
	"log"

	"github.com/anatolygg/tg_bot/internal/client/tg_bot"
	"github.com/anatolygg/tg_bot/internal/config"
	"github.com/anatolygg/tg_bot/internal/logger"
	"go.uber.org/zap"
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

	tgClient, err := tg_bot.New(token, log)
	if err != nil {
		log.Error("start bot failed", zap.Error(err))
	}

	tgClient.Start()
	// tgClient := telegram.New(token)

	// fetcher := fetcher.New()

	// processor := processor.New()
	// consumer.start(fetcher, processor)

}
