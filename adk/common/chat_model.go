package common

import (
	"context"
	"log"
	"os"

	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/model"
)

func NewChatModel(ctx context.Context) model.ToolCallingChatModel {

	//modelType := strings.ToLower(os.Getenv("MIMO_PRO"))
	cm, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		BaseURL: os.Getenv("MIMO_API_URL"),
		APIKey:  os.Getenv("MIMO_API_KEY"),
		Model:   os.Getenv("MIMO_PRO"),
	})
	if err != nil {
		log.Fatalf("openai.NewChatModel failed: %v", err)
	}
	return cm
}
