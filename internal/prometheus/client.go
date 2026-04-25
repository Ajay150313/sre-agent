package prometheus

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sravanthigorantla/sre-agent/pkg/models"
)

type Client struct {
	baseURL string
	client  *http.Client
}

func New(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		client:  &http.Client{Timeout: 10 * time.Second},
	}
}

type alertResponse struct {
	Data struct {
		Alerts []struct {
			Labels      map[string]string `json:"labels"`
			Annotations map[string]string `json:"annotations"`
			State       string            `json:"state"`
			ActiveAt    time.Time         `json:"activeAt"`
			Value       string            `json:"value"`
		} `json:"alerts"`
	} `json:"data"`
}

func (c *Client) GetActiveAlerts(ctx context.Context) ([]models.Alert, error) {
	url := fmt.Sprintf("%s/api/v1/alerts", c.baseURL)
	
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch alerts: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("prometheus returned status %d", resp.StatusCode)
	}

	var result alertResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode alerts: %w", err)
	}

	alerts := make([]models.Alert, 0, len(result.Data.Alerts))
	for _, a := range result.Data.Alerts {
		alerts = append(alerts, models.Alert{
			Name:        a.Labels["alertname"],
			Labels:      a.Labels,
			Annotations: a.Annotations,
			State:       a.State,
			ActiveAt:    a.ActiveAt,
		})
	}

	return alerts, nil
}

func (c *Client) QueryMetrics(ctx context.Context, query string) (float64, error) {
	url := fmt.Sprintf("%s/api/v1/query?query=%s", c.baseURL, query)
	
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return 0, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result struct {
		Data struct {
			Result []struct {
				Value []interface{} `json:"value"`
			} `json:"result"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	if len(result.Data.Result) > 0 && len(result.Data.Result[0].Value) > 1 {
		if v, ok := result.Data.Result[0].Value[1].(string); ok {
			var f float64
			fmt.Sscanf(v, "%f", &f)
			return f, nil
		}
	}

	return 0, fmt.Errorf("no data for query: %s", query)
}
