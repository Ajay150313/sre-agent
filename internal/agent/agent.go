package agent

type Agent struct {
	analyzer *Analyzer
}

func NewAgent(analyzer *Analyzer) *Agent {
	return &Agent{analyzer: analyzer}
}
