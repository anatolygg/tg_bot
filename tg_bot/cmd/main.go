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

	token := config.LoadEnv("TOKEN", "")

	tgClient, err := tg_bot.New(token, log)
	if err != nil {
		log.Error("create bot failed", zap.Error(err))
	}

	tgClient.NewML(cfg.MLService.URL)

	tgClient.Start()
}
