package ollama

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/ollama/ollama/api"
)

type Client struct {
	baseUrl  string
	httpClient *http.Client
}

func NewClient(host, port string) *Client {
	return &Client{
		baseUrl:    "http://" + host + ":" + port,
		httpClient:  &http.Client{},
	}
}

func (c *Client) ListModels(ctx context.Context) ([]string, error) {
	base, err := url.Parse(c.baseUrl)
	if err != nil {
		return nil, fmt.Errorf("invalid base client url: %w", err)
	}
	client := api.NewClient(base, c.httpClient)
	resp, err := client.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed getting llama models: %w", err)
	}
	models := make([]string, len(resp.Models))
	for i, model := range resp.Models{
		models[i] = model.Name
	}
	return models, nil
}

func (c *Client) Generate(ctx context.Context, prompt ,model string) (string, error) {
	base, err := url.Parse(c.baseUrl)
	if err != nil {
		return "", fmt.Errorf("invalid base client url: %w", err)
	}
	stream := false
	client := api.NewClient(base, c.httpClient)
	req := &api.GenerateRequest{
		Model: model,
		Prompt: prompt,
		Stream: &stream,
	}
	var response strings.Builder

	if err := client.Generate(ctx, req, func(gr api.GenerateResponse) error {
		response.WriteString(gr.Response)
		return nil
	}); err != nil {
		return "", fmt.Errorf("failed to generate response: %w", err)
	}
	return response.String(), nil
}

func (c *Client) HealthCheck(ctx context.Context) error {
	baseURL, err := url.Parse(c.baseUrl)
	if err != nil {
		return fmt.Errorf("invalid base URL: %w", err)
	}

	client := api.NewClient(baseURL, c.httpClient)

	_, err = client.List(ctx)
	if err != nil {
		return fmt.Errorf("health check failed: %w", err)
	}

	return nil
}

