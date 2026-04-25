package models

import "time"

type Severity string

const (
	SeverityCritical Severity = "critical"
	SeverityWarning  Severity = "warning"
	SeverityInfo     Severity = "info"
)

type Alert struct {
	Name        string            `json:"name"`
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
	State       string            `json:"state"`
	ActiveAt    time.Time         `json:"activeAt"`
	Value       float64           `json:"value"`
}

type Incident struct {
	ID          string    `json:"id"`
	Alert       Alert     `json:"alert"`
	Severity    Severity  `json:"severity"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	ResolvedAt  *time.Time `json:"resolvedAt,omitempty"`
	Diagnosis   string    `json:"diagnosis"`
	Remediation []string  `json:"remediation"`
	Logs        []string  `json:"logs"`
}

type AnalysisRequest struct {
	Alert   Alert              `json:"alert"`
	Logs    []string           `json:"logs"`
	Metrics map[string]float64 `json:"metrics"`
}
