package tg_bot

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	mlservice "github.com/anatolygg/tg_bot/internal/services/ml_service"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"go.uber.org/zap"
)

type TgBot struct {
	client     *bot.Bot
	ml_service *mlservice.MLModel
	logger     *zap.Logger
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

func (t *TgBot) NewML(url string) {
	t.ml_service = mlservice.New(url)
}

func (b *TgBot) Start() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	b.logger.Info("Запуск Telegram бота...")

	b.client.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, b.handleStart)
	b.client.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypePrefix, b.handleMessage)

	b.client.Start(ctx)
}

func (t *TgBot) handleMessage(ctx context.Context, b *bot.Bot, update *models.Update) {
	question := update.Message.Text
	answer := t.getAnswer(question)

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   answer,
	})
	if err != nil {

		t.logger.Error("Ошибка отправки сообщения", zap.Error(err))
	}
}

func (t *TgBot) handleStart(ctx context.Context, b *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	text := "Привет! Я бот для ответов на вопросы, связанные с МИФИ. Просто отправьте ваш вопрос, и я постараюсь помочь. Начнем?"

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   text,
	})
	if err != nil {
		fmt.Printf("Ошибка при отправке сообщения: %v", err)
	}
}

func (t *TgBot) getAnswer(question string) string {
	if question == "" {
		return "Задавайте свои вопросы!"
	}

	answer, err := t.ml_service.GetAnswer(question)
	if err != nil {
		t.logger.Error("Ошибка при запросе к ML-сервису", zap.Error(err))
		return "Произошла ошибка. Попробуйте позже."
	}

	return answer
}
