package chatbot

import (
	"testing"

	openai "github.com/sashabaranov/go-openai"
	"github.com/stretchr/testify/require"
)

func TestAddUserMessage(t *testing.T) {
	bot := NewChatBot("dummy-api-key")

	bot.AddUserMessage("Hello, world!")

	require.Contains(t, bot.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: "Hello, world!",
	})
}
