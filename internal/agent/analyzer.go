package agent

import (
	"context"
	"fmt"
	"time"

	"github.com/ajay150313/sre-agent/internal/llm"
	"github.com/ajay150313/sre-agent/internal/prometheus"
	"github.com/ajay150313/sre-agent/pkg/models"
)

type Analyzer struct {
	promClient *prometheus.Client
	llmClient  *llm.Client
}

func NewAnalyzer(promURL, openAIKey string) *Analyzer {
	return &Analyzer{
		promClient: prometheus.New(promURL),
		llmClient:  llm.New(openAIKey),
	}
}

func (a *Analyzer) AnalyzeAllAlerts(ctx context.Context) ([]models.Incident, error) {
	alerts, err := a.promClient.GetActiveAlerts(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch alerts: %w", err)
	}

	incidents := make([]models.Incident, 0, len(alerts))
	for _, alert := range alerts {
		incident, err := a.analyzeAlert(ctx, alert)
		if err != nil {
			continue
		}
		incidents = append(incidents, *incident)
	}

	return incidents, nil
}

func (a *Analyzer) analyzeAlert(ctx context.Context, alert models.Alert) (*models.Incident, error) {
	metrics := make(map[string]float64)
	
	if alert.Labels["job"] != "" {
		if cpu, err := a.promClient.QueryMetrics(ctx, 
			fmt.Sprintf("100 - (avg by(instance) (irate(node_cpu_seconds_total{mode=\"idle\",job=\"%s\"}[5m])) * 100)", alert.Labels["job"])); err == nil {
			metrics["cpu_usage"] = cpu
		}
	}

	analysis, err := a.llmClient.AnalyzeAlert(ctx, alert.Name, alert.Labels, []string{}, metrics)
	if err != nil {
		return nil, err
	}

	severity := models.SeverityWarning
	if analysis.Severity == "critical" {
		severity = models.SeverityCritical
	}

	return &models.Incident{
		ID:          fmt.Sprintf("INC-%d", time.Now().UnixNano()),
		Alert:       alert,
		Severity:    severity,
		Status:      "analyzed",
		CreatedAt:   time.Now(),
		Diagnosis:   analysis.RootCause,
		Remediation: analysis.Remediation,
		Logs:        []string{},
	}, nil
}
