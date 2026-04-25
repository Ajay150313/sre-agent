package llm

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

type Client struct {
	client *openai.Client
	model  string
}

type AnalysisResult struct {
	RootCause    string   `json:"root_cause"`
	Severity     string   `json:"severity"`
	Remediation  []string `json:"remediation"`
	Prevention   []string `json:"prevention"`
	Confidence   float64  `json:"confidence"`
}

func New(apiKey string) *Client {
	return &Client{
		client: openai.NewClient(apiKey),
		model:  openai.GPT4oMini,
	}
}

func (c *Client) AnalyzeAlert(ctx context.Context, alertName string, labels map[string]string, logs []string, metrics map[string]float64) (*AnalysisResult, error) {
	prompt := buildAnalysisPrompt(alertName, labels, logs, metrics)

	req := openai.ChatCompletionRequest{
		Model: c.model,
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleSystem, Content: SystemPrompt},
			{Role: openai.ChatMessageRoleUser, Content: prompt},
		},
		Temperature: 0.1,
		MaxTokens:   1000,
	}

	resp, err := c.client.CreateChatCompletion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("openai completion failed: %w", err)
	}

	if len(resp.Choices) == 0 {
		return nil, fmt.Errorf("no response from LLM")
	}

	content := resp.Choices[0].Message.Content
	var result AnalysisResult
	if err := json.Unmarshal([]byte(content), &result); err != nil {
		result = AnalysisResult{
			RootCause:   content,
			Severity:    "unknown",
			Remediation: []string{"Manual review required - LLM response unstructured"},
			Confidence:  0.5,
		}
	}

	return &result, nil
}

func (c *Client) EstimateCost(tokens int) float64 {
	return float64(tokens) * 0.0000002
}
