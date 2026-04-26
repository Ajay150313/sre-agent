# Changelog

## 0.1.0 - April 22, 2025

Initial release. SRE agent that reads Prometheus alerts and uses an LLM to figure out what's wrong and how to fix it.

Features:
- Connects to Prometheus for active alerts
- Uses OpenAI to analyze root causes
- Web dashboard for viewing results
- Docker support
- Health check endpoint

Coming next:
- Slack integration
- PagerDuty support
- Custom prompts for different alert types
- Better historical tracking
- Local LLM support (Ollama)
