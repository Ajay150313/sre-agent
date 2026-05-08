package llm

import "fmt"

const SystemPrompt = `You are an expert Site Reliability Engineer with 15 years of experience.
Your job is to analyze infrastructure alerts, logs, and metrics to provide:
1. Root cause analysis (2-3 sentences)
2. Severity assessment
3. Step-by-step remediation actions
4. Prevention recommendations

Be concise, technical, and actionable. Prioritize immediate mitigation over long-term fixes.`

func buildAnalysisPrompt(alertName string, labels map[string]string, logs []string, metrics map[string]float64) string {
	prompt := "ALERT: " + alertName + "\n\n"
	prompt += "LABELS:\n"
	for k, v := range labels {
		prompt += "- " + k + ": " + v + "\n"
	}
	
	if len(metrics) > 0 {
		prompt += "\nMETRICS:\n"
		for k, v := range metrics {
			prompt += "- " + k + ": " + fmt.Sprintf("%.2f", v) + "\n"
		}
	}
	
	if len(logs) > 0 {
		prompt += "\nRECENT LOGS:\n"
		for _, log := range logs {
			prompt += log + "\n"
		}
	}
	
	return prompt
}
