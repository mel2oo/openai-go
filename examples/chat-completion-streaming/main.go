package main

import (
	"context"

	"github.com/openai/openai-go"
)

func main() {
	client := openai.NewClient()

	ctx := context.Background()

	question := "选择一个你喜欢的词语"

	print("> ")
	println(question)
	println()

	stream := client.Chat.Completions.NewStreaming(ctx, openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(question),
		},
		Seed:         openai.Int(0),
		Model:        openai.ChatModelHN_R1,
		GuidedChoice: []string{"天气", "心情", "食物", "运动", "音乐"},
		// Reasoning: openai.Bool(true),
	})

	for stream.Next() {
		evt := stream.Current()
		if len(evt.Choices) > 0 {
			if len(evt.Choices[0].Delta.ReasoningContent) > 0 {
				print(evt.Choices[0].Delta.ReasoningContent)
			} else {
				print(evt.Choices[0].Delta.Content)
			}
		}
	}
	println()

	if err := stream.Err(); err != nil {
		panic(err.Error())
	}
}
