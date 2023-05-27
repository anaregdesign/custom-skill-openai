package openai

import (
	"context"
	"errors"
	"fmt"
	"github.com/anaregdesign/custom-skill-builder/model"
	"github.com/anaregdesign/go-aoai"
)

type Client struct {
	client *aoai.AzureOpenAI
}

func New(resourceName string, deproymentName string, apiVersion string, apiKey string) *Client {
	return &Client{
		client: aoai.New(resourceName, deproymentName, apiVersion, apiKey),
	}
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

func (c *Client) summarize(ctx context.Context, d *model.Data[string], languageCode string) (*model.Data[string], error) {
	system := fmt.Sprintf("You are the multilingual assistant. Please summarize given text with %s.", languageCode)
	summary, err := c.chatCompletion(ctx, system, d.Input, 0.1, 0.5)
	if err != nil {
		return nil, err
	}

	return &model.Data[string]{Output: summary}, nil
}

func (c *Client) SummarizeJA(ctx context.Context, d *model.Data[string]) (*model.Data[string], error) {
	return c.summarize(ctx, d, "ja")
}

func (c *Client) SummarizeEN(ctx context.Context, d *model.Data[string]) (*model.Data[string], error) {
	return c.summarize(ctx, d, "en")
}

func (c *Client) SummarizeZH(ctx context.Context, d *model.Data[string]) (*model.Data[string], error) {
	return c.summarize(ctx, d, "zh")
}

func (c *Client) SummarizeKO(ctx context.Context, d *model.Data[string]) (*model.Data[string], error) {
	return c.summarize(ctx, d, "ko")
}

func (c *Client) SummarizeFR(ctx context.Context, d *model.Data[string]) (*model.Data[string], error) {
	return c.summarize(ctx, d, "fr")
}

func (c *Client) SummarizeDE(ctx context.Context, d *model.Data[string]) (*model.Data[string], error) {
	return c.summarize(ctx, d, "de")
}

func (c *Client) SummarizeES(ctx context.Context, d *model.Data[string]) (*model.Data[string], error) {
	return c.summarize(ctx, d, "es")
}

func (c *Client) SummarizePT(ctx context.Context, d *model.Data[string]) (*model.Data[string], error) {
	return c.summarize(ctx, d, "pt")
}

func (c *Client) SummarizeIT(ctx context.Context, d *model.Data[string]) (*model.Data[string], error) {
	return c.summarize(ctx, d, "it")
}

func (c *Client) SummarizeNL(ctx context.Context, d *model.Data[string]) (*model.Data[string], error) {
	return c.summarize(ctx, d, "nl")
}

func (c *Client) SummarizePL(ctx context.Context, d *model.Data[string]) (*model.Data[string], error) {
	return c.summarize(ctx, d, "pl")
}

func (c *Client) SummarizeRU(ctx context.Context, d *model.Data[string]) (*model.Data[string], error) {
	return c.summarize(ctx, d, "ru")
}

func (c *Client) SummarizeTR(ctx context.Context, d *model.Data[string]) (*model.Data[string], error) {
	return c.summarize(ctx, d, "tr")
}

func (c *Client) SummarizeVI(ctx context.Context, d *model.Data[string]) (*model.Data[string], error) {
	return c.summarize(ctx, d, "vi")
}

func (c *Client) SummarizeAR(ctx context.Context, d *model.Data[string]) (*model.Data[string], error) {
	return c.summarize(ctx, d, "ar")
}

func (c *Client) SummarizeID(ctx context.Context, d *model.Data[string]) (*model.Data[string], error) {
	return c.summarize(ctx, d, "id")
}

func (c *Client) SummarizeTH(ctx context.Context, d *model.Data[string]) (*model.Data[string], error) {
	return c.summarize(ctx, d, "th")
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
