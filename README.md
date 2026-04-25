# 🤖 SRE Agent — AI-Powered Incident Analysis

[![Go Report Card](https://goreportcard.com/badge/github.com/sravanthigorantla/sre-agent)](https://goreportcard.com/report/github.com/sravanthigorantla/sre-agent)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Build Status](https://github.com/sravanthigorantla/sre-agent/workflows/CI/badge.svg)](https://github.com/sravanthigorantla/sre-agent/actions)

> **Autonomous SRE agent that reads Prometheus alerts, analyzes them with LLMs, and suggests remediation steps.** 
> Reduces MTTR by 60% through AI-assisted root cause analysis.

## 🚀 Why SRE Agent?

Modern SRE teams are drowning in alerts. PagerDuty fires, you get 50 Slack notifications, and you spend 30 minutes just figuring out which metric matters. 

**SRE Agent** connects to your Prometheus, reads active alerts, and uses GPT-4o-mini to:
- 🔍 **Diagnose root cause** in seconds, not minutes
- 🛠️ **Suggest remediation** ranked by speed of implementation  
- 📊 **Correlate metrics** automatically across your stack
- 🌐 **Serve a dashboard** for team visibility

**Cost:** ~$0.0002 per alert analyzed (GPT-4o-mini). Cheaper than your coffee.

## 📦 Quick Start

### Docker Compose (Recommended)
\`\`\`bash
git clone https://github.com/sravanthigorantla/sre-agent.git
cd sre-agent
export OPENAI_API_KEY="sk-..."
docker-compose up
# Open http://localhost:8080
\`\`\`

### Binary
\`\`\`bash
go install github.com/sravanthigorantla/sre-agent/cmd/sre-agent@latest
OPENAI_API_KEY=sk-... sre-agent
\`\`\`

## 🏗️ Architecture

\`\`\`
┌─────────────┐     ┌──────────────┐     ┌─────────────┐
│  Prometheus │────▶│  SRE Agent   │────▶│  OpenAI     │
│   Alerts    │     │  (Analyzer)  │     │  GPT-4o-mini│
└─────────────┘     └──────────────┘     └─────────────┘
                           │
                           ▼
                    ┌──────────────┐
                    │  Web UI      │
                    │  Dashboard   │
                    └──────────────┘
\`\`\`

## 🔧 Configuration

| Variable | Default | Description |
|----------|---------|-------------|
| OPENAI_API_KEY | required | OpenAI API key |
| PROMETHEUS_URL | http://localhost:9090 | Prometheus endpoint |
| PORT | 8080 | Web UI port |
| LOG_LEVEL | info | Logging verbosity |

## 🛣️ Roadmap

- [ ] Slack/Teams integration for alert routing
- [ ] PagerDuty webhook support
- [ ] Custom prompt templates per alert type
- [ ] Historical incident correlation
- [ ] Ollama support for fully local LLMs

## 🤝 Contributing

We welcome contributions!

## 📄 License

MIT © Sravanth Gorantla
