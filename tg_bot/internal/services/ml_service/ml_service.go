package mlservice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/anatolygg/tg_bot/internal/model"
)

type MLModel struct {
	url string
}

func New(url string) *MLModel {
	return &MLModel{url: url}
}

func (m *MLModel) GetAnswer(question string) (string, error) {
	requestBody, _ := json.Marshal(&model.MLRequest{Question: question})

	resp, err := http.Post(m.url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		slog.Error("err: ", slog.Any("err: ", err))
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("body: ", slog.Any("err: ", err))
		return "", fmt.Errorf("read answer failed: %w", err)
	}

	var result model.MLResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("unmarshal response failed: %w", err)
	}

	return result.Answer, nil
}
