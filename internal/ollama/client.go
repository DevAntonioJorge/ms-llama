package ollama

import "net/http"

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

func (c *Client) ListModels() ([]string, error) {
	// Implementation for listing models
	return nil, nil
}

func (c *Client) Generate(prompt string, model string) (string, error) {
	// Implementation for generating text
	return "", nil
}

func (c *Client) HealthCheck() (bool, error) {
	// Implementation for health check
	return true, nil
}