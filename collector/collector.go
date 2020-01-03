package collector

import "github.com/hexoul/metric-collector/generator"

// Collector gathers metric data
type Collector struct {
}

// New returns a new collector
func New(rpm uint) *Collector {
	metrics := []string{}

	if g, err := generator.New(metrics, makeRequestContext, rpm); err != nil {
		g.Start()
	}

	return &Collector{}
}

func makeRequestContext(metric string) {
	// temporal. it will be replaced
}
