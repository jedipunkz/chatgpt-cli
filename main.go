package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jedipunkz/chatgpt-cli/internal/chatbot"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	bot := chatbot.NewChatBot(os.Getenv("OPENAI_API_KEY"))

	for {
		fmt.Print("You: ")
		userInput, _ := reader.ReadString('\n')

		bot.AddUserMessage(userInput)

		aiResponse, err := bot.GenerateResponse()
		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			continue
		}

		fmt.Println("AI: ", aiResponse)
	}
}
