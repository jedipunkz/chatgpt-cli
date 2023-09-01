# Go ChatGPT Client

This project is a command-line interface (CLI) that allows continuous conversation with OpenAI's GPT-3 model, implemented in Go language using the [go-openai](https://github.com/sashabaranov/go-openai) library.

## Getting Started

First, you will need to install Go. Instructions for installation can be found [here](https://golang.org/doc/install).

Next, clone the repository:

```bash
git clone https://github.com/jedipunkz/chatgpt-cli.git
```

Then, navigate to the cloned directory:

```bash
cd chatgpt-cli
```

Before running the program, you need to set your OpenAI API key as an environment variable:

```bash
export OPENAI_API_KEY=your-api-key
```

Finally, you can run the program:

```bash
go run main.go -4   / /use openai.GPT4
go run main.go -3.5 // use openai.GPT3Dot5Turbo
```

## How to Use

After running the program, you will be prompted to type in a message. After you submit a message, the program will send it to the OpenAI API and print out the response from the AI.

## License

This project is licensed under the terms of the Apache License 2.0.

## Author

[@jedipunkz](https://github.com/jedipunkz)

