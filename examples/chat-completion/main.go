package main

import (
	"context"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func main() {
	client := openai.NewClient(
		option.WithBaseURL("http://10.20.152.76:8200/v1"),
	)

	ctx := context.Background()

	question := "Write me a haiku"

	print("> ")
	println(question)
	println()
	params := openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(question),
		},
		Seed:      openai.Int(0),
		Model:     openai.ChatModelHN_R1,
		Reasoning: openai.Bool(true),
	}

	completion, err := client.Chat.Completions.New(ctx, params)
	if err != nil {
		panic(err)
	}

	println(completion.Choices[0].Message.Content)
}
