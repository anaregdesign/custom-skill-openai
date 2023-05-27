package openai

import (
	"context"
	"errors"
	"github.com/anaregdesign/custom-skill-builder/model"
	aoai "github.com/anaregdesign/go-aoai"
)

type Client struct {
	client *aoai.AzureOpenAI
}

func (c *Client) Embed(ctx context.Context, d *model.Data[string]) (*model.Data[[]float64], error) {
	request := aoai.EmbeddingRequest{
		Inputs: []string{d.Input},
	}

	response, err := c.client.Embedding(ctx, request)
	if err != nil {
		return nil, err
	}

	return &model.Data[[]float64]{Output: response.Data[0].Embedding}, nil
}

func (c *Client) Summarize(ctx context.Context, d *model.Data[string]) (*model.Data[string], error) {
	system := "You are the multilingual speaker. Summarize given text with given language."
	summary, err := c.chatCompletion(ctx, system, d.Input, 0.1, 0.9)
	if err != nil {
		return nil, err
	}

	return &model.Data[string]{Output: summary}, nil
}

func (c *Client) chatCompletion(ctx context.Context, system string, prompt string, t float64, p float64) (string, error) {
	// https://platform.openai.com/docs/api-reference/chat/create
	request := aoai.ChatRequest{
		Messages: []aoai.ChatMessage{
			{
				Role:    "system",
				Content: system,
			},
			{
				Role:    "user",
				Content: prompt,
			},
		},
		Stream:      false,
		Temperature: t,
		TopP:        p,
	}

	response, err := c.client.ChatCompletion(ctx, request)
	if err != nil {
		return "", err
	}
	choice := response.Choices[0]

	if choice.FinishReason != "stop" {
		return "", errors.New("incomplete response: " + choice.FinishReason)
	}

	return response.Choices[0].Message.Content, nil

}
