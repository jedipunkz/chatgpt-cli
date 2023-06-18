package chatbot

import (
	"context"

	openai "github.com/sashabaranov/go-openai"
)

type ChatBot struct {
	client   *openai.Client
	messages []openai.ChatCompletionMessage
}

func NewChatBot(apiKey string) *ChatBot {
	return &ChatBot{
		client: openai.NewClient(apiKey),
	}
}

func (b *ChatBot) AddUserMessage(content string) {
	b.messages = append(b.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: content,
	})
}

func (b *ChatBot) GenerateResponse() (string, error) {
	resp, err := b.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       openai.GPT3Dot5Turbo,
			Messages:    b.messages,
			MaxTokens:   800,
			Temperature: 0.8,
		},
	)

	if err != nil {
		return "", err
	}

	aiResponse := resp.Choices[0].Message.Content
	b.messages = append(b.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: aiResponse,
	})

	return aiResponse, nil
}
