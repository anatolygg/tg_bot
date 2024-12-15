package tg_bot

import (
	"context"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"go.uber.org/zap"
)

type TgBot struct {
	client *bot.Bot
	logger *zap.Logger
}

func New(token string, logger *zap.Logger) (*TgBot, error) {
	client, err := bot.New(token)
	if err != nil {
		return nil, err
	}
	return &TgBot{
		client: client,
		logger: logger,
	}, nil
}

func (b *TgBot) Start() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	b.logger.Info("Запуск Telegram бота...")

	b.client.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypePrefix, handleMessage)

	b.client.Start(ctx)
}

func handleMessage(ctx context.Context, b *bot.Bot, update *models.Update) {
	question := update.Message.Text
	// b.logger.Info("получен вопрос", zap.String("question", question))

	// answer, err := b.mlService.GetAnswer(question)
	answer := "all good: " + question
	// if err != nil {
	// 	logger.Log.Errorf("Ошибка при запросе к ML: %v", err)
	// 	answer = "Произошла ошибка. Попробуйте позже."
	// }
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   answer,
	})
}
