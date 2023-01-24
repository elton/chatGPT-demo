package main

import (
	"context"
	"fmt"

	gogpt "github.com/sashabaranov/go-gpt3"
)

func main() {
	c := gogpt.NewClient("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")
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3Ada,
		MaxTokens: 5,
		Prompt:    "文",
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return
	}
	fmt.Printf("resp: %#v", resp)
}
