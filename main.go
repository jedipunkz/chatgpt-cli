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

	var version string

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-4":
			version = "4"
		case "-3.5":
			version = "3.5"
		default:
			fmt.Println("Invalid version specified! Use -4 or -3.5.")
			return
		}
	} else {
		fmt.Println("No version specified, defaulting to 4.")
		version = "4"
	}

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

		aiResponse, err := bot.GenerateResponse(version)
		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			continue
		}

		fmt.Println("AI: ", aiResponse)
	}
}
