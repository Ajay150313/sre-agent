# Get Started in 5 Minutes

## Docker Compose (Easiest)

```bash
git clone https://github.com/Ajay150313/sre-agent.git
cd sre-agent
export OPENAI_API_KEY="sk-your-key"
docker-compose up
```

Wait about 30 seconds. You should see "Analyzed X alerts" in the logs.

Open http://localhost:8080 and you'll see the dashboard.

## Just the Binary

If you want to run it locally without Docker:

```bash
go mod tidy
go run ./cmd/sre-agent/main.go
```

In another terminal:

```bash
curl http://localhost:8080/health
```

---

## Try It Out

Go to http://localhost:9090 (Prometheus). 

To trigger an alert for testing:

```bash
docker exec sre_agent-node-exporter-1 stress --cpu 2 --timeout 60s
```

Wait a couple minutes for the CPU alert to fire in Prometheus. Then click "Analyze Active Alerts" on the dashboard and watch it analyze the alert.

---

## Troubleshooting

**Getting "Connection refused" on port 8080?**
Check the logs: docker logs sre_agent-sre-agent-1

**Says OPENAI_API_KEY not found?**
Make sure you set it before running docker-compose: export OPENAI_API_KEY="sk-..."

**No alerts showing up?**
Check prometheus logs: docker logs sre_agent-prometheus-1

---

## What's Next

Check the README for architecture and how everything works. Look in examples for Prometheus configs. Want to help? See CONTRIBUTING.md.
