package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
)

func main() {
	now := time.Now()
	instruction := fmt.Sprintf("你是MiMo（中文名称也是MiMo），是小米公司研发的AI智能助手。\n今天的日期：%d年 %d月 %d日 ，你的知识截止日期是2024年12月。\n", now.Year(), now.Month(), now.Day())

	query := "用一句话解释 Eino 的 Component 设计解决了什么问题？"

	ctx := context.Background()
	chatModel, err := newChatModel(ctx)
	if err != nil {
		panic(err)
	}
	message := []*schema.Message{
		schema.SystemMessage(instruction),
		schema.UserMessage(query),
	}
	_, _ = fmt.Fprint(os.Stdout, "[assistant] ")
	sr, err := chatModel.Stream(ctx, message)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer sr.Close()

	for {
		frame, err := sr.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if frame != nil {
			_, _ = fmt.Fprint(os.Stdout, frame.Content)
		}
	}
	_, _ = fmt.Fprintln(os.Stdout)
}

func newChatModel(ctx context.Context) (model.ToolCallingChatModel, error) {

	return openai.NewChatModel(ctx, &openai.ChatModelConfig{
		BaseURL: os.Getenv("MIMO_API_URL"),
		APIKey:  os.Getenv("MIMO_API_KEY"),
		Model:   os.Getenv("MIMO_PRO"),
	})
}
