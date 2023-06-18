package main

import (
	"fmt"
	"os"

	"github.com/chzyer/readline"
	"github.com/jedipunkz/chatgpt-cli/internal/chatbot"
)

func main() {
	// to delete characters with Ctrl-H
	rl, err := readline.NewEx(&readline.Config{
		Prompt:          "You: ",
		HistoryFile:     "/tmp/readline.tmp",
		AutoComplete:    nil,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	bot := chatbot.NewChatBot(os.Getenv("OPENAI_API_KEY"))

	for {
		userInput, err := rl.Readline()
		if err == readline.ErrInterrupt {
			// User pressed Ctrl-C, exit program
			return
		} else if err != nil {
			fmt.Printf("Readline error: %v\n", err)
			continue
		}

		bot.AddUserMessage(userInput)

		aiResponse, err := bot.GenerateResponse()
		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			continue
		}

		fmt.Println("AI: ", aiResponse)
	}
}
